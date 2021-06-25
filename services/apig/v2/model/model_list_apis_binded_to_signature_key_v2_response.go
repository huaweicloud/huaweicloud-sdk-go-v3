package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApisBindedToSignatureKeyV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 本次查询返回的列表

	Bindings       *[]SignBindingApiResp `json:"bindings,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListApisBindedToSignatureKeyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisBindedToSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"ListApisBindedToSignatureKeyV2Response", string(data)}, " ")
}
