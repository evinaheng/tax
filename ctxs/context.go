package ctxs

// Context package

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
)

type contextKey int

// ContextCancelledError by user
const ContextCancelledError = "context canceled"

const (
	IPAddress contextKey = iota
	JSONContent
	QKey
	RouterVarContextKey
	TypeCodeKey
	NameKey
	PriceKey

	QQS        = "q"
	TypeCodeQS = "type_code"
	NameQS     = "name"
	PriceQS    = "price"
)

var (
	ContextQueryStringList = map[contextKey]string{
		QKey:        QQS,
		TypeCodeKey: TypeCodeQS,
		NameKey:     NameQS,
		PriceKey:    PriceQS,
	}
)

// Extract some value from *http.Request and append it into context
// If the value is not valid or there's any error, return the parent context as result
// add another exportable func to fetch the desired value from context
// make sure that req http.Request pointer is not nil
func (ck contextKey) getString(ctx context.Context) (string, bool) {
	str, ok := ctx.Value(ck).(string)
	return str, ok
}
func (ck contextKey) getInt(ctx context.Context) (int, bool) {
	str, _ := ck.getString(ctx)
	atoi, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	return atoi, true
}
func (ck contextKey) getInt64(ctx context.Context) (int64, bool) {
	str, _ := ck.getString(ctx)
	atoi, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, false
	}
	return atoi, true
}

// Save all values as string
// convert value into desired one on get context func
func AllQSToContext(ctx context.Context, req *http.Request) context.Context {
	for key, qs := range ContextQueryStringList {
		ctx = context.WithValue(ctx, key, req.FormValue(qs))
	}
	return ctx
}
func QFromContext(ctx context.Context) (string, bool) {
	return QKey.getString(ctx)
}

func NameKeyFromContext(ctx context.Context) (string, bool) {
	return NameKey.getString(ctx)
}

func TypeCodeKeyFromContext(ctx context.Context) (int, bool) {
	return TypeCodeKey.getInt(ctx)
}

func PriceKeyFromContext(ctx context.Context) (int, bool) {
	return PriceKey.getInt(ctx)
}

// ============
// IP ADDRESS
// ============
func IPAddressToContext(ctx context.Context, req *http.Request) context.Context {
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	return context.WithValue(ctx, IPAddress, net.ParseIP(ip))
}

func IPAddressFromContext(ctx context.Context) (net.IP, bool) {
	ip, ok := ctx.Value(IPAddress).(net.IP)
	return ip, ok
}

// JSONToContext value
func JSONToContext(ctx context.Context, req *http.Request) context.Context {

	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	return context.WithValue(ctx, JSONContent, body)
}

// JSONFromContext value
func JSONFromContext(ctx context.Context, i interface{}) bool {
	b, ok := ctx.Value(JSONContent).([]byte)
	if !ok {
		return false
	}

	if err := json.Unmarshal(b, i); err != nil {
		log.Println("func JSONFromContext", err)
		return false
	}
	return true
}

// Function to fetch context from *http.Request
// set timeout if QS timeout is parseable into time.Duration
// implement the save value to context funcs as optional parameter
func GetContextFromRequest(req *http.Request, opt ...func(context.Context, *http.Request) context.Context) context.Context {
	ctx := req.Context()

	for _, f := range opt {
		ctx = f(ctx, req)
	}

	return ctx
}

func GetAllContextFromRequest(req *http.Request) context.Context {
	contentType := req.Header.Get("Content-Type")
	if contentType == "application/vnd.api+json" || contentType == "application/json" {
		return GetContextFromRequest(req, AllQSToContext, JSONToContext)
	}
	return GetContextFromRequest(req, AllQSToContext)
}
