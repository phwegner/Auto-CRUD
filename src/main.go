package main
import (
	"fmt"
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"bytes"

)
type Data struct{
	id string
}




func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fmt.Fprintf(w, "key: %s\n", vars["id"])


	buf, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			log.Print("bodyErr ", bodyErr.Error())
			http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
			return
		}
		
		log.Printf(buf)

	return 
}

// Existing code from above
func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    //myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/create", create).Methods("POST")
   
    log.Fatal(http.ListenAndServe(":9009", myRouter))
}



func main() {
	fmt.Println("Rest API v1.0 - ##--Auto-CRUD--##")

	// Route handles & endpoints
	http.HandleFunc("/create", create)
	

	// Start server
	handleRequests()
}