package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificatePrivateKeyEchoResponse Response Object
type ShowCertificatePrivateKeyEchoResponse struct {

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**：ELB证书的私钥回显开关。  **取值范围**： - true：开启回显，证书的列表、详情、创建和更新接口响应字段将会显示私钥内容。 - false：关闭回显，证书的所有接口响应不显私钥内容，使用脱敏内容（*****）代替。
	PrivateKeyEcho *bool `json:"private_key_echo,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowCertificatePrivateKeyEchoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificatePrivateKeyEchoResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificatePrivateKeyEchoResponse", string(data)}, " ")
}
