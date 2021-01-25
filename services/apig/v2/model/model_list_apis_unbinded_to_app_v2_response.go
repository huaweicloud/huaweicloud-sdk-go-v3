package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApisUnbindedToAppV2Response struct {
	// 符合条件的API总数
	Total *int32 `json:"total,omitempty"`
	// 本次返回的列表长度
	Size *int32 `json:"size,omitempty"`
	// 本次返回的API列表
	Apis           *[]AppAuthUnBindedApiResp `json:"apis,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListApisUnbindedToAppV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisUnbindedToAppV2Response struct{}"
	}

	return strings.Join([]string{"ListApisUnbindedToAppV2Response", string(data)}, " ")
}
