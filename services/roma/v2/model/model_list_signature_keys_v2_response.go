package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSignatureKeysV2Response Response Object
type ListSignatureKeysV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的签名密钥列表
	Signs          *[]SignatureWithBindNum `json:"signs,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSignatureKeysV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSignatureKeysV2Response struct{}"
	}

	return strings.Join([]string{"ListSignatureKeysV2Response", string(data)}, " ")
}
