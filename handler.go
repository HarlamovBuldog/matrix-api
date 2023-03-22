package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func GetMatrixFromReq(r *http.Request) (Matrix, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	matrix, err := CreateFromCSV(records)
	if err != nil {
		return nil, err
	}
	return matrix, nil
}

// echo godoc
// @Summary Print matrix in readable format.
// @Description Return the matrix as a string in matrix format.
// @ID echo
// @Success 200 all ok
// @Failure 400 bad input
// @Router /echo [get]
func echo(w http.ResponseWriter, r *http.Request) {
	matrix, err := GetMatrixFromReq(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("bad request: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, matrix.Echo(), "\n")
}

// invert godoc
// @Summary Invert matrix and prints it in readable format.
// @Description Return the matrix as a string in matrix format where the columns and rows are inverted
// @ID invert
// @Success 200 all ok
// @Failure 400 bad input
// @Router /invert [get]
func invert(w http.ResponseWriter, r *http.Request) {
	matrix, err := GetMatrixFromReq(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("bad request: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, matrix.Invert().Echo(), "\n")
}

// flatten godoc
// @Summary Prints matrix in one line.
// @Description Return the matrix as a 1 line string, with values separated by commas.
// @ID flatten
// @Success 200 all ok
// @Failure 400 bad input
// @Router /flatten [get]
func flatten(w http.ResponseWriter, r *http.Request) {
	matrix, err := GetMatrixFromReq(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("bad request: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, matrix.Flatten(), "\n")
}

// sum godoc
// @Summary Return the sum of the integers in the matrix.
// @Description Return the sum of the integers in the matrix.
// @ID sum
// @Success 200 all ok
// @Failure 400 bad input
// @Router /sum [get]
func sum(w http.ResponseWriter, r *http.Request) {
	matrix, err := GetMatrixFromReq(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("bad request: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, matrix.Sum(), "\n")
}

// multiply godoc
// @Summary Return the product of the integers in the matrix
// @Description Return the product of the integers in the matrix
// @ID multiply
// @Success 200 all ok
// @Failure 400 bad input
// @Router /multiply [get]
func multiply(w http.ResponseWriter, r *http.Request) {
	matrix, err := GetMatrixFromReq(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("bad request: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, matrix.Mul(), "\n")
}
