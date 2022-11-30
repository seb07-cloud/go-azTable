package manipulateAzTable

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

type apiError struct {
	Err    string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func (t Table) MakeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJson(w, e.Status, e)
				return
			}
			writeJson(w, http.StatusInternalServerError, apiError{Err: "internal server", Status: http.StatusInternalServerError})
		}
	}
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func (t Table) GetHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Method == http.MethodGet {
		t.PartitionKey = r.URL.Query().Get("PartitionKey")
		t.RowKey = r.URL.Query().Get("RowKey")

		message, err := t.Get()
		if err != nil {
			return apiError{Err: "couldnt get value", Status: http.StatusBadRequest}
		}
		
		return writeJson(w, http.StatusOK, message)

	} else {

		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}

	}
}

func (t Table) GetSingleHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Method == http.MethodGet {
		t.PartitionKey = r.URL.Query().Get("PartitionKey")
		t.RowKey = r.URL.Query().Get("RowKey")
		t.PropertyName = r.URL.Query().Get("PropertyName")

		message, err := t.GetSingle()
		if err != nil {
			return apiError{Err: "couldnt get value", Status: http.StatusBadRequest}
		}

		return writeJson(w, http.StatusOK, message)

	} else {

		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}

	}
}

func (t Table) UpdateHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Method == http.MethodPost {

		r.ParseForm()

		data := struct {
			Method      string
			URL         *url.URL
			Submissions url.Values
		}{
			r.Method,
			r.URL,
			r.Form,
		}

		s := data.Submissions
		t.PartitionKey = string(s.Get("PartitionKey"))
		t.RowKey = string(s.Get("RowKey"))
		t.PropertyName = string(s.Get("PropertyName"))
		t.PropertyValue = string(s.Get("PropertyValue"))

		if t.ValidateParams(t.PropertyName) && t.ValidateParams(t.PropertyValue) {
			message, err := t.Update()
			if err != nil {
				return apiError{Err: "couldnt get value", Status: http.StatusBadRequest}
			}
			return writeJson(w, http.StatusOK, message)

		} else {

			return apiError{Err: "not enough parameters", Status: http.StatusBadRequest}

		}

	} else {

		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}

	}
}

func (t Table) DeleteHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Method == http.MethodDelete {

		r.ParseForm()

		data := struct {
			Method      string
			URL         *url.URL
			Submissions url.Values
		}{
			r.Method,
			r.URL,
			r.Form,
		}

		s := data.Submissions
		t.PropertyName = fmt.Sprintln(string(s.Get("PropertyName")))

		if reflect.ValueOf(t.PropertyName).IsValid() {

			fmt.Print(w, "Here a delete property function should be implemented")

		} else {

			http.Error(w, "not enough parameters.", http.StatusBadRequest)

		}

	} else {

		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}

	}

	return writeJson(w, http.StatusOK, "ok")
}