package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSignatureKeysBindedToApiV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size" xml:"size"`

	// 满足条件的记录数
	Total int64 `json:"total" xml:"total"`

	// API与签名密钥的绑定关系列表
	Bindings       *[]SignApiBindingInfo `json:"bindings,omitempty" xml:"bindings"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListSignatureKeysBindedToApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSignatureKeysBindedToApiV2Response struct{}"
	}

	return strings.Join([]string{"ListSignatureKeysBindedToApiV2Response", string(data)}, " ")
}
