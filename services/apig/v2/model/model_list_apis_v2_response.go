package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApisV2Response struct {
	// 满足条件的分组总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回的列表长度

	Size *int32 `json:"size,omitempty"`
	// 本次查询到的API列表

	Apis           *[]ApiInfoPerPage `json:"apis,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListApisV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisV2Response struct{}"
	}

	return strings.Join([]string{"ListApisV2Response", string(data)}, " ")
}
