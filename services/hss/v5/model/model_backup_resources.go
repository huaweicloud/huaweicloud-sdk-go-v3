package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupResources **参数解释**: 开启备份功能新版参数，必填；若为空代表兼容之前绑定HSS_projectid的存储库 **约束限制**: 不涉及 **取值范围**: 取值0-20个BackupResources对象 **默认取值**: 不涉及
type BackupResources struct {

	// **参数解释**: 选择需要绑定的存储库ID，不为空 **约束限制**: 不涉及 **取值范围**: 字符长度0-64 **默认取值**: 不涉及
	VaultId *string `json:"vault_id,omitempty"`

	// **参数解释**: 需要开启备份功能的主机情况列表 **约束限制**: 不涉及 **取值范围**: 取值0-20条主机id **默认取值**: 不涉及
	ResourceList *[]ResourceInfo `json:"resource_list,omitempty"`
}

func (o BackupResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupResources struct{}"
	}

	return strings.Join([]string{"BackupResources", string(data)}, " ")
}
