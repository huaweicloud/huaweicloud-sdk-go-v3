package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSignatureKeysBindedToApiV2Response struct {
	// 本次查询满足条件的总数

	Total *int32 `json:"total,omitempty"`
	// 本次查询返回的列表长度

	Size *int32 `json:"size,omitempty"`
	// 本次查询返回的列表

	Bindings       *[]SignBindingApiResp `json:"bindings,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListSignatureKeysBindedToApiV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSignatureKeysBindedToApiV2Response struct{}"
	}

	return strings.Join([]string{"ListSignatureKeysBindedToApiV2Response", string(data)}, " ")
}
