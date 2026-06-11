package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CertInfoOption **参数解释：** CCM证书信息。 **取值范围：** 开启ssl，不传此参数，表示使用默认证书。
type CertInfoOption struct {

	// **参数解释：** 证书ID。 **取值范围：** 根据CCM证书列表接口获取证书ID。
	CertId string `json:"cert_id"`

	// **参数解释：** 证书类型。 **取值范围：**   - PCA：CCM PCA 证书。   - SSL：CCM SSL 证书。
	CertType string `json:"cert_type"`
}

func (o CertInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertInfoOption struct{}"
	}

	return strings.Join([]string{"CertInfoOption", string(data)}, " ")
}
