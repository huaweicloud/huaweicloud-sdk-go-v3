package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateCompositeHostsResponse Response Object
type MigrateCompositeHostsResponse struct {

	// host_id列表
	HostIds *[]string `json:"host_ids,omitempty"`

	// 策略ID（目标企业项目下的策略ID）
	PolicyId *string `json:"policy_id,omitempty"`

	// 证书ID（目标企业项目下的证书ID）
	CertificateId  *string `json:"certificate_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateCompositeHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateCompositeHostsResponse struct{}"
	}

	return strings.Join([]string{"MigrateCompositeHostsResponse", string(data)}, " ")
}
