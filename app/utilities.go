package app

import (
	"encoding/json"
	"net/http"
)

func (a *App) WriteJSON(w http.ResponseWriter, status int, data interface{}) error {

	js, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return err
		}
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return nil
	}

	return nil
}

func (a *App) ErrorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}
	theError := jsonError{
		Message: err.Error(),
	}
	_ = a.WriteJSON(w, http.StatusBadRequest, theError)
	if err != nil {

	}
}
func (a *App) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return nil
	}

	return nil
}
