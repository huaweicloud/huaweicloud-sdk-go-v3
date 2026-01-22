package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDomainVerificationResponse Response Object
type CheckDomainVerificationResponse struct {

	// 主域名
	Domain *string `json:"domain,omitempty"`

	// 校验值，解析值或者文件内容
	VerifyContent *string `json:"verify_content,omitempty"`

	// 验证结果，true为验证成功确认归属，false为归属情况未确认
	VerifyResult   *bool `json:"verify_result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckDomainVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDomainVerificationResponse struct{}"
	}

	return strings.Join([]string{"CheckDomainVerificationResponse", string(data)}, " ")
}
