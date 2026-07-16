package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificateAuthority struct {

	// **参数解释：** Base64编码的CCE集群根证书 **约束限制：** 该字段仅在获取指定的集群时返回 **取值范围：** 不涉及 **默认取值：** 不涉及
	Data *string `json:"data,omitempty"`
}

func (o CertificateAuthority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateAuthority struct{}"
	}

	return strings.Join([]string{"CertificateAuthority", string(data)}, " ")
}
