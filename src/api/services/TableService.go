package services

import (
	"CRUD/src/api/data"
	"CRUD/src/api/utilities"
	"github.com/gorilla/mux"
	"net/http"
)

func FindAllTables(w http.ResponseWriter, r *http.Request) {
	var resp Response
	tables, err := models.FindAllTables()

	if err != nil {
		err = utilities.ErrorJSON(w, err, http.StatusBadRequest)
		if err != nil {
			return
		}
		return
	}

	resp.Data = tables
	resp.Message = "SUCCESS"
	err = utilities.WriteJSON(w, http.StatusOK, &resp)

	if err != nil {
		return
	}

}

func FindTableById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, castErr := utilities.CastFromStringToInt(params["id"], w)
	if castErr == true {
		return
	}

	var resp Response
	table, err := models.FindTableById(id)

	if err != nil {
		err = utilities.ErrorJSON(w, err, http.StatusBadRequest)
		if err != nil {
			return
		}
		return
	}

	resp.Data = table
	resp.Message = "SUCCESS"
	jsonErr := utilities.WriteJSON(w, http.StatusOK, &resp)
	if jsonErr != nil {
		return
	}
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	var table = models.GetTable()
	var resp Response

	err := utilities.ReadJSON(w, r, &table)
	if err != nil {
		jsonErr := utilities.ErrorJSON(w, err, http.StatusBadRequest)
		if jsonErr != nil {
			return
		}
		return
	}

	table, err = models.CreateTable(table)
	if err != nil {
		return
	}

	resp.Data = table
	resp.Message = "SUCCESS"
	jsonErr := utilities.WriteJSON(w, http.StatusOK, &resp)
	if jsonErr != nil {
		return
	}
}

func UpdateTable(w http.ResponseWriter, r *http.Request) {
	var table = models.GetTable()
	var resp Response

	err := utilities.ReadJSON(w, r, &table)
	if err != nil {
		jsonErr := utilities.ErrorJSON(w, err, http.StatusBadRequest)
		if jsonErr != nil {
			return
		}
		return
	}

	table, err = models.UpdateTable(table)
	if err != nil {
		return
	}

	resp.Data = table
	resp.Message = "SUCCESS"
	jsonErr := utilities.WriteJSON(w, http.StatusOK, &resp)
	if jsonErr != nil {
		return
	}
}

func DeleteTableById(w http.ResponseWriter, r *http.Request) {
	var resp Response
	params := mux.Vars(r)

	id, castErr := utilities.CastFromStringToInt(params["id"], w)
	if castErr == true {
		return
	}

	err := models.DeleteTableById(id)
	if err != nil {
		jsonErr := utilities.ErrorJSON(w, err, http.StatusBadRequest)
		if jsonErr != nil {
			return
		}
		return
	}

	resp.Message = "SUCCESS"
	jsonErr := utilities.WriteJSON(w, http.StatusOK, &resp)
	if jsonErr != nil {
		return
	}
}
