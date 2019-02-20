package entity

type Bills struct {
	ID      int     `json:"id" db:"id"`
	Name    string  `json:"name" db:"name"`
	TaxCode int     `json:"tax_code" db:"tax_code"`
	Price   float64 `json:"price" db:"price"`
}

type FormattedBills struct {
	Name       string  `json:"name"`
	TaxCode    int     `json:"tax_code"`
	Type       string  `json:"type"`
	Refundable bool    `json:"is_refundable"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	Amount     float64 `json:"amount"`
}

type BillsInfo struct {
	GrandTotal  float64 `json:"grand_total"`
	SubTotal    float64 `json:"sub_total"`
	TaxSubTotal float64 `json:"tax_sub_total"`
}
