package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgreementVo struct {

	// 协议id
	Id *string `json:"id,omitempty"`

	// 协议名称
	AgreementName *string `json:"agreement_name,omitempty"`

	// 协议类型
	AgreementType *string `json:"agreement_type,omitempty"`

	// 协议类型名称
	AgreementTypeName *string `json:"agreement_type_name,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 版本
	Version *int32 `json:"version,omitempty"`
}

func (o AgreementVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgreementVo struct{}"
	}

	return strings.Join([]string{"AgreementVo", string(data)}, " ")
}
