package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
   "log"
	"github.com/gorilla/mux"
)
type MOVIE struct{
	Id string  `json:"id"`
	Name string  `json:"name"`
	Director *Director  `json:"director"`
}
type Director struct{
	FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}


func allMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r);
	if len(params)==0 {
      errors.New("NO PARAMS SPECIFIED BITCH")
			return;
	}
	for _,x := range movie{
		if x.Id == params["id"]{
			json.NewEncoder(w).Encode(x)
			return
		}
	}
}

func createFunc(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-Type", "application/json")
	 var newMovie MOVIE;
	 json.NewDecoder(r.Body).Decode(&newMovie)
	 newMovie.Id=strconv.Itoa(rand.Intn(10000000))
	 movie=append(movie,newMovie)
	 json.NewEncoder(w).Encode(movie)

}

func updateFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newMovie MOVIE;
	params:=mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.Id = strconv.Itoa(rand.Intn(100000))
  for i,j := range movie{
     if j.Id == params["id"] {
			movie= append (movie[:i],movie[i+1:]...)
			movie=append(movie,newMovie)
			break
		 }
	}
   json.NewEncoder(w).Encode(movie)
}


func deleteFunc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r);
	if len(params)==0 {
      errors.New("NO PARAMS SPECIFIED BITCH")
			return;
	}
	for i,x := range movie{
		if x.Id == params["id"]{
		movie=	append(movie[:i],movie[i+1:]...)
			return
		}
	}
}


var movie []MOVIE
func main(){
	 // movie =append(movie,MOVIE{Id : "2343" , Name : "scary movie", Director :&Director{"bi","k"}})
      r:=mux.NewRouter();
			r.HandleFunc("/movies",allMovies).Methods(("GET"))
			r.HandleFunc("/movies/{id}",getMovie).Methods(("GET"))
			r.HandleFunc("/movies",createFunc).Methods(("POST"))
			r.HandleFunc("/movies/{id}",deleteFunc).Methods(("DELETE"))
			r.HandleFunc("/movies/{id}",updateFunc).Methods(("PUT"))
      fmt.Printf("starting server at port 8000 \n")
			log.Fatal(http.ListenAndServe(":8000",r))
}