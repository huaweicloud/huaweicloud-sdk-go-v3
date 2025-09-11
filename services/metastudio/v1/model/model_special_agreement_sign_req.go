package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecialAgreementSignReq 增加租户权限
type SpecialAgreementSignReq struct {

	// 当前服务协议类型: AUTO_PAY: 自动支付协议
	AgreementType string `json:"agreement_type"`

	// 是否为去签署,true:签署;false:取消签署
	ToSign bool `json:"to_sign"`
}

func (o SpecialAgreementSignReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecialAgreementSignReq struct{}"
	}

	return strings.Join([]string{"SpecialAgreementSignReq", string(data)}, " ")
}
