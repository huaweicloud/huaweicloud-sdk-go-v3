package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCitiesResponse struct {
	// 查询个数，成功的时候返回。

	Count *int32 `json:"count,omitempty"`
	// 城市信息列表，成功的时候返回，具体参见表2。

	Cities         *[]City `json:"cities,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCitiesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCitiesResponse struct{}"
	}

	return strings.Join([]string{"ListCitiesResponse", string(data)}, " ")
}
