package db

import "gorm.io/gorm"

func RunSeeds(db *gorm.DB) error {
	return seedProducts(db)
}

func seedProducts(db *gorm.DB) error {
	// This just makes things slightly easier, though I am not a big fan of this approach.
	products := []Product{
		{
			Id:               "P001",
			Name:             "Frozen Prawns 16/20",
			Species:          "Prawn",
			OriginCountry:    "India",
			ProductionMethod: "aquaculture",
			PricePerTonne:    1850,
			SupplierName:     "BlueWave Foods",
			Certification:    "ASC",
		},
		{
			Id:               "P002",
			Name:             "Atlantic Cod Fillet",
			Species:          "Cod",
			OriginCountry:    "Iceland",
			ProductionMethod: "wild",
			PricePerTonne:    2150,
			SupplierName:     "Nordic Sea Co",
			Certification:    "MSC",
		},
		{
			Id:               "P003",
			Name:             "Smoked Salmon Slices",
			Species:          "Salmon",
			OriginCountry:    "Scotland",
			ProductionMethod: "aquaculture",
			PricePerTonne:    3200,
			SupplierName:     "Highland Smokehouse",
			Certification:    "ASC",
		},
		{
			Id:               "P004",
			Name:             "Haddock Portions",
			Species:          "Haddock",
			OriginCountry:    "Norway",
			ProductionMethod: "wild",
			PricePerTonne:    2050,
			SupplierName:     "Fjord Fisheries",
			Certification:    "MSC",
		},
		{
			Id:               "P005",
			Name:             "Canned Tuna CHUNK",
			Species:          "Tuna",
			OriginCountry:    "Thailand",
			ProductionMethod: "wild",
			PricePerTonne:    1450,
			SupplierName:     "Siam Oceanic",
			Certification:    "None",
		},
		{
			Id:               "P006",
			Name:             "Mussels Vac-Pack",
			Species:          "Mussel",
			OriginCountry:    "Netherlands",
			ProductionMethod: "aquaculture",
			PricePerTonne:    950,
			SupplierName:     "Delta Shellfish",
			Certification:    "ASC",
		},
	}

	return db.Create(products).Error
}
