package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSignatureKeysV2Response struct {
	// 符合条件的签名密钥的数量
	Total *int32 `json:"total,omitempty"`
	// 本次查询到的签名密钥列表
	Signs *[]SignatureCreateResp `json:"signs,omitempty"`
	// 本次返回的列表长度
	Size           *int32 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSignatureKeysV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSignatureKeysV2Response struct{}"
	}

	return strings.Join([]string{"ListSignatureKeysV2Response", string(data)}, " ")
}
