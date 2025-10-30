package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupResp struct {

	// **参数解释** 卷备份ID **取值范围** 字符长度0-65535位
	Id *string `json:"id,omitempty"`

	// **参数解释** 磁盘备份名称 **取值范围** 字符长度0-65535位
	Name *string `json:"name,omitempty"`

	// **参数解释** 备份类型 **取值范围** 字符长度0-65535位
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释** 备份所在的存储库ID **取值范围** 字符长度0-65535位
	VaultId *string `json:"vault_id,omitempty"`

	// **参数解释** 备份状态 **取值范围** 字符长度0-65535位
	Status *string `json:"status,omitempty"`

	// **参数解释** 创建时间 **取值范围** 取值0-2147483647
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// **参数解释** 资源ID 对应host主机ID **取值范围** 字符长度0-65535位
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释** 资源类型 **取值范围** 字符长度0-65535位
	ResourceType *string `json:"resource_type,omitempty"`

	// **参数解释** 资源名称 对应host主机名称 **取值范围** 字符长度0-65535位
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o BackupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupResp struct{}"
	}

	return strings.Join([]string{"BackupResp", string(data)}, " ")
}
