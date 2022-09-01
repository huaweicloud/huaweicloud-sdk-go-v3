package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceBackupPolicy struct {

	// 备份策略ID
	BackupPolicyId *string `json:"backup_policy_id,omitempty" xml:"backup_policy_id"`

	// 创建时间。格式为：2022-04-11T09:45:24.790Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间。格式为：2022-04-12T02:22:03.269Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	Policy *BackupPolicy `json:"policy,omitempty" xml:"policy"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`
}

func (o InstanceBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceBackupPolicy struct{}"
	}

	return strings.Join([]string{"InstanceBackupPolicy", string(data)}, " ")
}
