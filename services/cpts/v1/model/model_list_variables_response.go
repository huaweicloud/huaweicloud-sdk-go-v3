package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListVariablesResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
	// variable_list

	VariableList   *[]ListVariablesResqVariableList `json:"variable_list,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListVariablesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVariablesResponse struct{}"
	}

	return strings.Join([]string{"ListVariablesResponse", string(data)}, " ")
}