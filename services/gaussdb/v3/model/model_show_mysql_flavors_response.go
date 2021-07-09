package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlFlavorsResponse struct {
	// 实例规格信息列表

	Flavors        *[]MysqlFlavorsInfo `json:"flavors,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowMysqlFlavorsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlFlavorsResponse", string(data)}, " ")
}
