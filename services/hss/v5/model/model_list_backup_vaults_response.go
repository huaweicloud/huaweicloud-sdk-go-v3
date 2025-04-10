package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupVaultsResponse Response Object
type ListBackupVaultsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 查询备份存储库列表
	DataList       *[]BackupVaultInfo `json:"data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBackupVaultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupVaultsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupVaultsResponse", string(data)}, " ")
}
