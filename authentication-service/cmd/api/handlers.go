package main

import "net/http"

func (app *Application) authenticate(w http.ResponseWriter, _ *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "authentication service is running",
	}

	err := app.WriteJson(w, http.StatusOK, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
