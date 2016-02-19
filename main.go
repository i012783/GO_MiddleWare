package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/codegangsta/negroni"
)

func SteveHandler(b bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if b!=false {
			fmt.Fprintf(w, "true")
			return
		}
		fmt.Fprintf(w, "false")
	}
}

func IsAuthenticated(token string) negroni.Handler  {
	au := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		fmt.Fprintf(w, "1.  Build logic to validate token -> UAA/Check_Token endpoint\n")

				if token != "VALID"{
					fmt.Fprintf(w, "     Not a valid token -> redirect user to /login and exit chain\n\n\n")
				}else{
					fmt.Fprintf(w, "2.  Process Authorization Next\n")
					fmt.Fprintf(w, "-------------------------------------------------------------\n\n\n")
					next(w,r)}
			}
	return negroni.HandlerFunc(au)
}

func IsAuthorized(AccessQuery string) negroni.Handler  {
	az := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Fprintf(w, "3.  ACS endpoint Entitlement Query:  %s\n",  AccessQuery)
		fmt.Fprintf(w, "4.  Add logic and build HTTP Rest client to call ACS\n\n")
		next(w,r)
	}
	return negroni.HandlerFunc(az)
}


func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/", SteveHandler(true))
	mux.PathPrefix("/protected").Methods("GET").Handler(negroni.New(negroni.Handler(IsAuthenticated("VALID")), negroni.Handler(IsAuthorized("http://avi.acs.io/engine/update"))))

	log.Println("Listening....")
	http.ListenAndServe(":8080", mux)
}