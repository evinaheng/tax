package constant

const (
	FOOD          = 1
	TOBACCO       = 2
	ENTERTAINMENT = 3
)

var TaxMap = map[int]string{
	FOOD:          "Food",
	TOBACCO:       "Tobacco",
	ENTERTAINMENT: "Entertainment",
}

var ItemRefund = map[int]bool{
	FOOD:          true,
	TOBACCO:       false,
	ENTERTAINMENT: false,
}
