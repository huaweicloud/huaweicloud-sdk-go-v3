package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleInstance struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 实例类型。支持集群、副本集、以及单节点。 取值   - Sharding   - ReplicaSet   - Single
	Mode *string `json:"mode,omitempty"`

	Datastore *RecycleDatastore `json:"datastore,omitempty"`

	// 计费方式。 - 取值为“0”，表示按需计费。 - 取值为“1”，表示包年/包月计费。
	PayMode *string `json:"pay_mode,omitempty"`

	// 企业项目ID，取值为“0”，表示为default企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 备份ID
	BackupId *string `json:"backup_id,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 删除时间
	DeletedAt *string `json:"deleted_at,omitempty"`

	// 保留截止时间
	RetainedUntil *string `json:"retained_until,omitempty"`
}

func (o RecycleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstance struct{}"
	}

	return strings.Join([]string{"RecycleInstance", string(data)}, " ")
}
