package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct : base de los datos que pediremos a nuestra API

type Base struct {
	ID     int    `json:"ID"`
	Nombre string `json:"Nombre"`
	Genero string `json:"Genero"`
}

type Peticiones []Base

var peticion = Peticiones{{

	ID:     1,
	Nombre: "juan",
	Genero: "Masculino",
}}

//funcion Index

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenidos a mi API REST")
}

//Funcion para crear registro en nuestra api

func C_Peticion(w http.ResponseWriter, r *http.Request) {
	//Creamos la variable de guardar datos
	var N_Peticion Base
	//Creamos una variable temporal para la validacion
	reqBody, err := ioutil.ReadAll(r.Body)

	//Evaluamos si se encuentra algun error en la peticion

	if err != nil {
		fmt.Fprintf(w, "Los datos ingresados no son correctos")
	}
	//Asignacion de datos que son leidos por la variable body
	json.Unmarshal(reqBody, &N_Peticion)
	//Incremento de nuestro ID

	N_Peticion.ID = len(peticion) + 1

	//añadimos el registro
	peticion = append(peticion, N_Peticion)

	//Indicamos los tipos de datos que se ingresaran en nuestro servicio
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(N_Peticion)
}

//Funcion para eliminar todos los datos 

func F_Eliminar(w http.ResponseWriter, r * http.Request){

	eliminacion:= mux.Vars(r)

	BusquedaID, err := strconv.Atoi(eliminacion["id"])

	if err != nil{
		return
	}

	for i, t := range peticion {
		if t.ID == BusquedaID{
			peticion = append(peticion[:i], peticion[i+1:]...)

			fmt.Fprintf(w, "El registro %v se ha eliminado correctamente de la API",BusquedaID)
		}
	}
}



// Funcion para mostrar
func M_Todo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peticion)
}

//Rutas

func main() {
	ruta := mux.NewRouter().StrictSlash(true)

	// Ruta index

	ruta.HandleFunc("/", Index)

	//ruta para todo el contenido

	ruta.HandleFunc("/Datos", M_Todo).Methods("GET")

	//Ruta para crear peticiones en la Api
	ruta.HandleFunc("/Datos", C_Peticion).Methods("POST")

	// Ruta para eliminar datos de la API
     ruta.HandleFunc("/Datos/{id}", F_Eliminar).Methods("DELETE")

	//Inicializacion de servidor

	log.Fatal(http.ListenAndServe(":8000", ruta))

}
