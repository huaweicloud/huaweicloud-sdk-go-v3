package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 生成证书的安全模式： true：安全模式（证书由系统托管，每个用户名只能生成一个证书，每个组织生成上限100个） false：非安全模式（证书由用户自己保障，不限制生成数量）
type CreateBlockchainCertByUserNameRequestBody struct {

	// 生成证书的安全模式： true：安全模式（证书由系统托管，每个用户名只能生成一个证书，每个组织生成上限100个） false：非安全模式（证书由用户自己保障，不限制生成数量）
	SecurityMode *bool `json:"security_mode,omitempty" xml:"security_mode"`
}

func (o CreateBlockchainCertByUserNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBlockchainCertByUserNameRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBlockchainCertByUserNameRequestBody", string(data)}, " ")
}
