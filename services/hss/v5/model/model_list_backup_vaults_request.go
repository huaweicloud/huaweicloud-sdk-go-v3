package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupVaultsRequest Request Object
type ListBackupVaultsRequest struct {

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// 备份存储库名称
	VaultName *string `json:"vault_name,omitempty"`

	// 备份存储库id
	VaultId *string `json:"vault_id,omitempty"`
}

func (o ListBackupVaultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupVaultsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupVaultsRequest", string(data)}, " ")
}
