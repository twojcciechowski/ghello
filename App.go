package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var signKey=[]byte("WnVRO[J0s~i4uTbv&/kA((?qoY4mnDjg-=g~w3(Sij<C@[e%K%Yr3L8VnB&>HSs_j%!B[2cE(t1poZifDg34ed61>-N.")

func main() {

	ge:=os.Getenv("GET_ENDPOINT")
	pe:=os.Getenv("POST_ENDPOINT")

	router:=mux.NewRouter()

	nonAuthRoute:=router.NewRoute().Subrouter()
	nonAuthRoute.HandleFunc("/newJwt", handlerNewJwt).Methods(http.MethodGet)
	if len(ge)>0 {
		nonAuthRoute.HandleFunc("/"+ge, handlerGet).Methods(http.MethodGet)
	}
	if len(pe)>0 {
		nonAuthRoute.HandleFunc("/"+pe, handlerPost).Methods(http.MethodPost)
	}

	authRoute:=router.NewRoute().Subrouter()
	authRoute.Use(jwtAuth)

	if len(ge)>0 {
		authRoute.HandleFunc("/jwt/"+ge, handlerGet).Methods(http.MethodGet)
	}
	if len(pe)>0 {
		authRoute.HandleFunc("/jwt/"+pe, handlerPost).Methods(http.MethodPost)
	}

	port:=os.Getenv("PORT")
	if len(port)==0 {
		port="8080"
	}
	
	s:=http.Server{
		Addr:              "0.0.0.0:"+port,
		Handler:           router,
	}
	log.Fatal(s.ListenAndServe())
}

func handlerNewJwt(w http.ResponseWriter, r *http.Request){
	claims:=&jwt.StandardClaims{
		ExpiresAt: time.Now().Add(100*time.Hour).Unix(),
		NotBefore: time.Now().Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	stoken,err:=token.SignedString(signKey)
	if err!=nil {
		log.Printf("Error: %v",err)
		_,_=fmt.Fprintf(w,"Error: %s",err.Error())
		w.WriteHeader(500)
	}

	_,e:=io.WriteString(w, stoken)
	if e!=nil {
		_,_=fmt.Fprint(w,e.Error())
	}

}

func jwtAuth(http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		tknStr := r.Header.Get("Authorization")
		log.Print("Authorization header: ",tknStr)

		tknStr = strings.Replace(tknStr, "Bearer ", "", 1)
		tknStr = strings.Replace(tknStr, "bearer ", "", 1)

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
			return signKey, nil
		})

		if err==nil {
			log.Printf("Decoded token: %+v",token.Claims)
		}

		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusUnauthorized)

			switch err {
			case jwt.ErrSignatureInvalid:
				_,_=fmt.Fprint(w,"incorrect signature")
			default:
				_,_=fmt.Fprint(w,"incorrect jwt token")
			}
			
			return
		}

		//Block 'none' algorithm
		if token.Method.Alg() == jwt.SigningMethodNone.Alg() {
			_,_=fmt.Fprint(w,"'none' algorithm is not supported")
			return
		}
	})
}


func handlerGet(w http.ResponseWriter, r *http.Request){

	_,e:=io.WriteString(w, os.Getenv("GET_MSG"))
	if e!=nil {
		_,_=fmt.Fprint(w,e.Error())
	}

}

func handlerPost(w http.ResponseWriter, r *http.Request){
	_,e:=io.WriteString(w, os.Getenv("POST_MSG"))
	if e!=nil {
		_,_=fmt.Fprint(w,e.Error())
	}
}
