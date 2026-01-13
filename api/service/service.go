package service

import (
	"encoding/json"
	"log"
	"net/http"
	"oo-api/db"
	"oo-api/models"
)

type Service struct {
	db *db.DB
}

func (s *Service) writeErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Write([]byte(message))
	w.WriteHeader(status)
}

// New returns a new instance of a service.
func New(dbConn *db.DB) *Service {
	return &Service{db: dbConn}
}

func (s *Service) GetProducts(w http.ResponseWriter, req *http.Request) {
	dbProducts, err := s.db.GetAllProducts()
	if err != nil {
		log.Printf("could not load the products: %s", err)
		s.writeErrorResponse(w, "could not load the products", http.StatusInternalServerError)
		return
	}

	products := []models.Product{}
	for _, dbProduct := range dbProducts {
		products = append(products, models.Product{
			Id:               dbProduct.Id,
			Name:             dbProduct.Name,
			OriginCountry:    dbProduct.OriginCountry,
			ProductionMethod: dbProduct.ProductionMethod,
			Species:          dbProduct.Species,
			SupplierName:     dbProduct.SupplierName,
			PricePerTonne:    dbProduct.PricePerTonne,
			Certification:    dbProduct.Certification,
		})
	}

	response, err := json.Marshal(products)
	if err != nil {
		log.Printf("could not load the products: %s", err)
		s.writeErrorResponse(w, "could not load the products", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
