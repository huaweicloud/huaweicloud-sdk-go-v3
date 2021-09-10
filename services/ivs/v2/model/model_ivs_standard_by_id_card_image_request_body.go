package model

import (
	"encoding/json"

	"strings"
)

type IvsStandardByIdCardImageRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsStandardByIdCardImageRequestBodyData `json:"data"`
}

func (o IvsStandardByIdCardImageRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IvsStandardByIdCardImageRequestBody struct{}"
	}

	return strings.Join([]string{"IvsStandardByIdCardImageRequestBody", string(data)}, " ")
}
