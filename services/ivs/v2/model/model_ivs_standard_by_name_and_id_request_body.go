package model

import (
	"encoding/json"

	"strings"
)

type IvsStandardByNameAndIdRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsStandardByNameAndIdRequestBodyData `json:"data"`
}

func (o IvsStandardByNameAndIdRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IvsStandardByNameAndIdRequestBody struct{}"
	}

	return strings.Join([]string{"IvsStandardByNameAndIdRequestBody", string(data)}, " ")
}
