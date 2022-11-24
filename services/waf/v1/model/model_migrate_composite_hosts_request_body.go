package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迁移域名请求体
type MigrateCompositeHostsRequestBody struct {

	// host_id列表
	HostIds []string `json:"host_ids"`

	// 策略ID（目标企业项目下的策略ID）
	PolicyId string `json:"policy_id"`

	// 证书ID（目标企业项目下的证书ID）
	CertificateId *string `json:"certificate_id,omitempty"`
}

func (o MigrateCompositeHostsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateCompositeHostsRequestBody struct{}"
	}

	return strings.Join([]string{"MigrateCompositeHostsRequestBody", string(data)}, " ")
}
