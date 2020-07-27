package main

import (
	"github.com/go-macaron/binding"
	log "github.com/sirupsen/logrus"
	"gopkg.in/macaron.v1"
	"labs/macaron-binding/model"
	"net/http"
	"strings"
)

func main() {
	m := macaron.Classic()
	m.Post("/test", binding.Bind(model.TestRequest{}), func(w http.ResponseWriter, mRequest model.TestRequest) {
		log.Info(mRequest)
		w.WriteHeader(http.StatusOK)
	})
	m.Get("/test", binding.Bind(model.TestGetRequest{}), func(w http.ResponseWriter, mRequest model.TestGetRequest) {
		log.Info(mRequest)
		w.WriteHeader(http.StatusOK)
	})

	binding.AddRule(&binding.Rule{
		IsMatch: func(rule string) bool {
			return rule == "GmailValidation"
		},
		IsValid: func(errs binding.Errors, name string, v interface{}) (bool, binding.Errors) {
			fieldSplit := strings.Split(v.(string), "@")
			if strings.ToLower(fieldSplit[1]) != "gmail.com" {
				errs = append(errs, binding.Error{
					FieldNames:     []string{name},
					Classification: "GmailValidation",
					Message:        "email is not a gmail",
				})
				return false, errs
			}
			return true, errs
		},
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", m))
}
