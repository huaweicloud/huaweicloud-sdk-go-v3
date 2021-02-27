package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProvincesResponse struct {
	// 查询个数，成功的时候返回。
	Count *int32 `json:"count,omitempty"`
	// 省份信息列表，成功的时候返回，具体参见表3。
	Provinces      *[]Province `json:"provinces,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListProvincesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProvincesResponse struct{}"
	}

	return strings.Join([]string{"ListProvincesResponse", string(data)}, " ")
}
