package utilities

import (
	"net/http"
	"strconv"
)

func CastFromStringToInt(value string, w http.ResponseWriter) (int, bool) {
	intVar, err := strconv.Atoi(value)

	if err != nil {
		jsonErr := ErrorJSON(w, err, http.StatusBadRequest)
		if jsonErr != nil {
			return 0, true
		}
		return 0, true
	}

	return intVar, false
}
