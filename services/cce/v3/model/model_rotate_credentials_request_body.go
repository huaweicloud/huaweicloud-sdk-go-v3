package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RotateCredentialsRequestBody struct {

	// **参数解释：**  需要轮转的组件名称。 **约束限制：** 不涉及 **取值范围：** - all：轮转CCE集群证书。 - service-account-controller：轮转ServiceAccount token相关证书。 - custom：轮转用户自有证书，指定此参数时，需同时指定certContent参数。  **默认取值：** 不涉及
	Component string `json:"component"`

	// **参数解释：**  轮转证书后，用于验证ServiceAccount Token签名的旧证书保留时间。  为了保证基于旧证书签发的ServiceAccount Token在证书轮转后能验签通过，CCE会保留老证书一段时间，具体规则如下： - 首次轮转时，CCE会保留创建集群时生成的证书； - 从第二次轮转开始，CCE会保留老证书一段时间，默认24小时。用户可以通过当前参数配置保留的时间。  **约束限制：** 不涉及 **取值范围：** 0-8784（小时） **默认取值：** 24（小时）
	CertificateExpirationTime *int32 `json:"certificateExpirationTime,omitempty"`

	CertContent *AuthenticatingProxy `json:"certContent,omitempty"`
}

func (o RotateCredentialsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateCredentialsRequestBody struct{}"
	}

	return strings.Join([]string{"RotateCredentialsRequestBody", string(data)}, " ")
}
