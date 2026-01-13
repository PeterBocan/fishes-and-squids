package models

type Product struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	Species          string  `json:"species"`
	OriginCountry    string  `json:"origin_country"`
	ProductionMethod string  `json:"production_method"`
	PricePerTonne    float64 `json:"price_per_tonne"`
	SupplierName     string  `json:"supplier_name"`
	Certification    string  `json:"certification"`
}
