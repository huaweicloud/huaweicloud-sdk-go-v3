package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupVaultsRequest Request Object
type ListBackupVaultsRequest struct {

	// 偏移量：指定返回记录的开始位置
	Offset int32 `json:"offset"`

	// 每页显示个数
	Limit int32 `json:"limit"`

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
