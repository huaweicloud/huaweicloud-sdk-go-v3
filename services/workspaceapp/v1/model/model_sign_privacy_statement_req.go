package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignPrivacyStatementReq 签署隐私声明请求。
type SignPrivacyStatementReq struct {

	// 隐私声明版本号。
	Version string `json:"version"`
}

func (o SignPrivacyStatementReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignPrivacyStatementReq struct{}"
	}

	return strings.Join([]string{"SignPrivacyStatementReq", string(data)}, " ")
}
