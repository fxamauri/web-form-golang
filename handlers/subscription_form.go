package handlers

import (
	"fmt"
	"net/http"

	"github.com/fxamauri/web-form-golang/controllers"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "erro ao fazer parse do form: %v", err)
			return
		}

		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Fprintf(w, "erro ao fazer parse do form: %v", err)
			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/subscription_form.html")
}
