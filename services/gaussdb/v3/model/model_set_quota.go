package model

import (
	"encoding/json"

	"strings"
)

type SetQuota struct {
	// 企业项目ID。

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// 实例个数配额。

	InstanceQuota int32 `json:"instance_quota"`
	// CPU核数配额。

	VcpusQuota int32 `json:"vcpus_quota"`
	// 内存使用配额，单位为GB。

	RamQuota int32 `json:"ram_quota"`
}

func (o SetQuota) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetQuota struct{}"
	}

	return strings.Join([]string{"SetQuota", string(data)}, " ")
}
