package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type User struct {

	// **参数解释**： 客户端证书。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClientCertificateData *string `json:"client-certificate-data,omitempty"`

	// **参数解释**： 包含来自TLS客户端密钥文件的PEM编码数据。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ClientKeyData *string `json:"client-key-data,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}
