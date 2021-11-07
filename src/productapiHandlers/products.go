package productapiHandlers

import (
	"data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// Handler
type Products struct {
	l *log.Logger
}

// Constructor
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entrypoint for the handler and satisfies http.handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle request to add a product (POST)
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("More than 1 ID")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("More than 1 capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("unable to convert to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if id == 9223372036854775807 {
			p.l.Println("ID out of range")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, rw, r)
		return
	}

	//catch all
	//if no method is satisfied
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// sub handler that returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to get Json", http.StatusInternalServerError)
	}
}

// sub handler that creates/replaces a product from the data store
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Parse Json", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	//expect the ID in the URI
	p.l.Println("Handle PUT")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Parse Json", http.StatusBadRequest)
		return
	}

	//p.l.Println("got id", id)
	err2 := data.UpdateProduct(id, prod)
	if err2 != nil {
		p.l.Printf("Could not update record", id, err2)
		http.Error(rw, "Error Updating record", http.StatusInternalServerError)
		return
	}
}
