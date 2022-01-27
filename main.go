package main //incluir módulo main

import (
	"encoding/json" //Importar librerías
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/dgrijalva/jwt-go" //importar repositorios
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*")) //Carpeta de plantillas HTML

type Empleado struct { //Generar estructura Empleado con datos de correo y contraseña
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
}

var jwtRev []byte = []byte("GraphQL") // Generar variable jwtRev para la clave secreta

var cuenta []Empleado = []Empleado{ //Se generan registros de empleados
	Empleado{
		Correo:     "Lorenzo",
		Contrasena: "lord_",
	},
	Empleado{
		Correo:     "Juan",
		Contrasena: "King",
	},
}

var cuentaTipo *graphql.Object = graphql.NewObject(graphql.ObjectConfig{ //Generación de servidor GRAPHQL
	Name: "cuenta",
	Fields: graphql.Fields{
		"correo": &graphql.Field{
			Type: graphql.String,
		},
		"contrasena": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func main() { //Función main
	http.HandleFunc("/", Index)       //Ejecutar servicio HTTP en página principal
	http.HandleFunc("/login", Token)  //Ejecutar servicio HTTP en página login, utilizando la función Token
	http.ListenAndServe(":8080", nil) //Ejecutar servicio HTTP puerto 8080
}

func Index(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "index", nil) //redirigir a pagina principal index
}

func ValidarJWT(t string) (interface{}, error) { //Función de ValidarJWT
	if t == "" {
		return nil, errors.New("authorizacion de Token")
	}
	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error de")
		}
		return jwtRev, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //Sección de reclamaciones
		var decodeToken interface{}
		mapstructure.Decode(claims, &decodeToken)
		return decodeToken, nil
	} else {
		return nil, errors.New("token inválido")
	}
}

func Token(rw http.ResponseWriter, r *http.Request) { //Función Token
	var usuario Empleado
	_ = json.NewDecoder(r.Body).Decode(&usuario)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //generación de toKen, seccion Header
		"correo":     usuario.Correo, //Generación de claims o carga
		"contrasena": usuario.Contrasena,
	})
	tokenResult, error := token.SignedString(jwtRev) //firma del token
	if error != nil {
		fmt.Println(error)
	}
	rw.Header().Set("Content-type", "application/json")
	rw.Write([]byte(tokenResult)) //Impresión del tóken
}
