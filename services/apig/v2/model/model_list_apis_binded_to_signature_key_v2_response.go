package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApisBindedToSignatureKeyV2Response struct {
	// 本次查询满足条件的总数

	Total *int32 `json:"total,omitempty"`
	// 本次查询返回的列表长度

	Size *int32 `json:"size,omitempty"`
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
