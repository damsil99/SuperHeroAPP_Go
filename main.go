package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/***************************************************************************
* Estructuras
***************************************************************************/

// Biography: biografía de un héroe.
type Biography struct {
	FullName string `json:"fullName,omitempty"`
}

// PowerStats: estadísticas de poder de un héroe.
type PowerStats struct {
	Intelligence int `json:"intelligence,omitempty"`
	Strength     int `json:"strength,omitempty"`
	Speed        int `json:"speed,omitempty"`
	Durability   int `json:"durability,omitempty"`
	Power        int `json:"power,omitempty"`
	Combat       int `json:"combat,omitempty"`
}

// Images: imágenes de un héroe en diferentes tamaños.
type Images struct {
	Xs string `json:"xs,omitempty"`
	Sm string `json:"sm,omitempty"`
	Md string `json:"md,omitempty"`
	Lg string `json:"lg,omitempty"`
}

// Heroes: agrupa toda la información relevante de un héroe.
type Heroes struct {
	Name       string      `json:"name,omitempty"`
	Biography  *Biography  `json:"biography,omitempty"`
	PowerStats *PowerStats `json:"powerstats,omitempty"`
	Images     *Images     `json:"images,omitempty"`
}

/***************************************************************************
* Base de datos
***************************************************************************/
var BaseDatosHeroes []Heroes

/***************************************************************************
* Controla las solicitudes GET.
***************************************************************************/
func EndPointHero(w http.ResponseWriter, req *http.Request) {

	// Extrae las variables de la ruta desde la solicitud HTTP.
	// Obtiene el valor de "name" que se envía como parte de la URL.
	params := mux.Vars(req)

	for _, item := range BaseDatosHeroes {
		if item.Name == params["name"] {
			// Codifica y devuelve el héroe que coincide con el nombre en el parámetro de la URL.
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// Si no se encuentra el héroe, devuelve un estado 404 y un mensaje de error.
	w.WriteHeader(http.StatusNotFound) //404 Not Found
	w.Write([]byte("EL HÉROE NO ESTÁ EN LA BASE DE DATOS :("))

}

// Funcion principal del programa
func main() {

	// Crea un nuevo router de mux para manejar las rutas.
	router := mux.NewRouter()

	// Llena la "base de datos".

	// ----------------------------------------------------
	BaseDatosHeroes = append(BaseDatosHeroes, Heroes{
		Name: "Wolverine",
		Biography: &Biography{
			FullName: "John Logan"},
		PowerStats: &PowerStats{
			Intelligence: 63,
			Strength:     32,
			Speed:        50,
			Durability:   100,
			Power:        89,
			Combat:       100},
		Images: &Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg"}})

	// ----------------------------------------------------
	BaseDatosHeroes = append(BaseDatosHeroes, Heroes{
		Name: "Spider-Man",
		Biography: &Biography{
			FullName: "Peter Parker"},
		PowerStats: &PowerStats{
			Intelligence: 90,
			Strength:     55,
			Speed:        67,
			Durability:   75,
			Power:        74,
			Combat:       85},
		Images: &Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg"}})

	// ----------------------------------------------------
	BaseDatosHeroes = append(BaseDatosHeroes, Heroes{
		Name: "Iron Man",
		Biography: &Biography{
			FullName: "Tony Stark"},
		PowerStats: &PowerStats{
			Intelligence: 100,
			Strength:     85,
			Speed:        58,
			Durability:   85,
			Power:        100,
			Combat:       64},
		Images: &Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg"}})

	// ----------------------------------------------------
	BaseDatosHeroes = append(BaseDatosHeroes, Heroes{
		Name: "Black Widow",
		Biography: &Biography{
			FullName: "Natasha Romanoff"},
		PowerStats: &PowerStats{
			Intelligence: 75,
			Strength:     13,
			Speed:        33,
			Durability:   30,
			Power:        36,
			Combat:       100},
		Images: &Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg"}})

	// ----------------------------------------------------
	BaseDatosHeroes = append(BaseDatosHeroes, Heroes{
		Name: "Thor",
		Biography: &Biography{
			FullName: "Thor Odinson"},
		PowerStats: &PowerStats{
			Intelligence: 69,
			Strength:     100,
			Speed:        83,
			Durability:   100,
			Power:        100,
			Combat:       100},
		Images: &Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg"}})

	// Define la ruta y el controlador para manejar la solicitud GET de superhéroes.
	router.HandleFunc("/api/superhero", EndPointHero).Queries("hero", "{name}").Methods("GET")

	// Inicia el servidor en el puerto 8080 utilizando el router configurado.
	log.Fatal(http.ListenAndServe(":8080", router))

}
