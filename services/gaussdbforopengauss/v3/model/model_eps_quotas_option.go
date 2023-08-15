package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EpsQuotasOption struct {

	// 企业项目Id。
	EnterpriseProjectsId string `json:"enterprise_projects_id"`

	// 实例的配额。取值范围：实际创建的实例个数 ~ 100,000。
	InstanceQuota *int32 `json:"instance_quota,omitempty"`

	// cpu的配额。取值范围：实际使用的cpu核数 ~ 2,147,483,646。
	VcpusQuota *int32 `json:"vcpus_quota,omitempty"`

	// 内存的配额。单位GB。取值范围：实际使用的内存 ~ 2,147,483,646。
	RamQuota *int32 `json:"ram_quota,omitempty"`

	// 存储空间的配额。单位：GB。取值范围：实际使用的存储空间 ~ 2,147,483,646。
	VolumeQuota *int32 `json:"volume_quota,omitempty"`
}

func (o EpsQuotasOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpsQuotasOption struct{}"
	}

	return strings.Join([]string{"EpsQuotasOption", string(data)}, " ")
}
