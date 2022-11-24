package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据恢复请求体
type RestoreBackupReq struct {

	// 执行策略（true：全部覆盖，false：全部跳过，默认为true）
	Overwrite *bool `json:"overwrite,omitempty"`

	// 目标文件夹
	TargetFolder *string `json:"target_folder,omitempty"`

	// 目标项目ID
	TargetProjectId string `json:"target_project_id"`
}

func (o RestoreBackupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBackupReq struct{}"
	}

	return strings.Join([]string{"RestoreBackupReq", string(data)}, " ")
}
