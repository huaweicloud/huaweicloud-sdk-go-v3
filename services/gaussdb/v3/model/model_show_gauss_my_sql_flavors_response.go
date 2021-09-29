package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGaussMySqlFlavorsResponse struct {
	// 实例规格信息列表

	Flavors        *[]MysqlFlavorsInfo `json:"flavors,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowGaussMySqlFlavorsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlFlavorsResponse", string(data)}, " ")
}
