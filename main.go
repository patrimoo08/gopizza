// TODO -
// Validate Post inputs
// Auto Generate OrderNumber
// Improve Toppings Situation - Validate against a list of toppings?
// Subclass toppings for light, normal, extra?
// Rate Limit HTTP Post requests by IP - https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
//

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pizza struct {
	OrderNumber string  `json:"OrderNumber"`
	Toppings    string  `json:"Toppings"`
	Crust       float64 `json:"Crust"`
	Sauce       float64 `json:"Sauce"`
	ExtraCheese bool    `json:"ExtraCheese"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Pizzas []Pizza

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Papa Milhouse's Pizza API! Demo for VON!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPizzas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPizzas")
	json.NewEncoder(w).Encode(Pizzas)
}

func returnSinglePizza(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["OrderNumber"]

	fmt.Fprintf(w, "Key: "+key)
	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, Pizza := range Pizzas {
		if Pizza.OrderNumber == key {
			json.NewEncoder(w).Encode(Pizza)
		}
	}
}

func createNewPizza(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var pizza Pizza
	json.Unmarshal(reqBody, &pizza)
	// update our global Articles array to include
	// our new Pizza
	Pizzas = append(Pizzas, pizza)

	json.NewEncoder(w).Encode(pizza)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/Pizzas", returnAllPizzas)
	myRouter.HandleFunc("/Pizza", createNewPizza).Methods("POST")
	myRouter.HandleFunc("/Pizza/{OrderNumber}", returnSinglePizza)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	// Pizza DB
	Pizzas = []Pizza{
		Pizza{OrderNumber: "1", Toppings: "Pepperoni, Jalapenos, Garlic", Crust: 2, Sauce: 1, ExtraCheese: true},
		Pizza{OrderNumber: "2", Toppings: "Chicken, Bacon, Onion", Crust: 1, Sauce: 2, ExtraCheese: false},
	}
	handleRequests()

}
