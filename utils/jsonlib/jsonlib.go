// Package jsonlib
package jsonlib

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// Serializer a interface to string, return json string if successful.
// Return "",error if json.Marshal() fail
func Serializer(value interface{}) (string, error) {
	byte, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	byteData, err := SerializerToBytes(string(byte))
	if err != nil {
		return "", err
	}
	return string(byteData), nil
}

// SerializerToBytes a interface to byte array, return json byte array if successful.
// Return nil,error if json.Marshal() fail
func SerializerToBytes(value interface{}) ([]byte, error) {
	byte, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return byte, nil
}

// Deserializer a string to interface, result is object if successful.
// Return error if json.Unmarshal() fail
func Deserializer(value string, result interface{}) error {
	return DeserializerFromBytes([]byte(value), result)
}

// DeserializerFromBytes a byte array to interface, return json string if successful.
// Return "",error if json.Marshal() fail
func DeserializerFromBytes(value []byte, result interface{}) error {
	if json.Valid(value) {
		return json.Unmarshal(value, result)
	}
	return errors.New("invalid json byte array")
}
