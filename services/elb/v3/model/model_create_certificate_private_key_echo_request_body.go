package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificatePrivateKeyEchoRequestBody This is a auto create Body Object
type CreateCertificatePrivateKeyEchoRequestBody struct {

	// **参数解释**：ELB证书的私钥回显开关。  **约束限制**：不涉及  **取值范围**： - true：开启回显，证书的列表、详情、创建和更新接口响应字段将会显示私钥内容。 - false：关闭回显，证书的所有接口响应不显私钥内容，使用脱敏内容（*****）代替。  **默认取值**：不涉及
	PrivateKeyEcho bool `json:"private_key_echo"`
}

func (o CreateCertificatePrivateKeyEchoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificatePrivateKeyEchoRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificatePrivateKeyEchoRequestBody", string(data)}, " ")
}
