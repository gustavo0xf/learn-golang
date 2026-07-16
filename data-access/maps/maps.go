package main

import "errors"

type Map map[string]string

var (
	errNotFound      = errors.New("[!] unable to find the key on the map")
	errAlreadyExists = errors.New("[!] key already exists")
)

func (m Map) searchMap(key string) (string, error) {
	found, ok := m[key]

	if !ok {
		return "", errNotFound
	}

	return found, nil
}

func (m Map) updateMap(key, val string) (string, error) {
	_, err := m.searchMap(key)

	if err == errNotFound {
		m[key] = val
		return "[*] map updated!", nil
	} else {
		return "", errAlreadyExists
	}
}

func main() {
	myMap := Map{"uuid": "89869fd7-7eaa-463f-987a-f224a7c85d36"}
	myMap.searchMap("uuid")
}
