package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tracker-cli/types"
)

var (
	TRACKER_CORE_BASE_ROUTE = "http://127.0.0.1:8080"
	ErrConvertingJsonToType = errors.New("Error Convirtiendo")
	ErrRequestingApi        = errors.New("Error consultando")
	ErrReadingBody          = errors.New("Error leyendo")
)

func mergeBaseRoute(next string) string {
	return fmt.Sprintf("%s/%s", TRACKER_CORE_BASE_ROUTE, next)
}

func GetAllTransactions() *map[string]types.Transaction {
	path := mergeBaseRoute("transaction?message=transactionKeeperGetAllTransactions")

	t, err := baseRequestGet[map[string]types.Transaction](path)
	if err != nil {
		log.Fatalln(err)
	}

	return t
}

func AddTransaction(t *types.Transaction) error {
	path := mergeBaseRoute("transaction?message=transactionKeeperAddTransactions")
	_, err := baseRequestPost[interface{}](path, t)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func baseRequestGet[Type any](path string) (*Type, error) {
	r, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, ErrReadingBody
	}

	var t Type
	err = json.Unmarshal(s, &t)
	if err != nil {
		return nil, ErrConvertingJsonToType
	}

	return &t, nil
}

func baseRequestPost[Type any](path string, body interface{}) (*Type, error) {
	contentType := "application/json"

	r, err := json.Marshal(body)
	if err != nil {
		return nil, ErrConvertingJsonToType
	}

	buffer := bytes.NewBuffer(r)
	response, err := http.Post(path, contentType, buffer)
	if err != nil {
		log.Fatalln(err)
		return nil, ErrRequestingApi
	}
	defer response.Body.Close()

	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, ErrReadingBody
	}

	var t Type
	err = json.Unmarshal(s, &t)
	if err != nil {
		return nil, ErrConvertingJsonToType
	}

	return &t, nil
}
