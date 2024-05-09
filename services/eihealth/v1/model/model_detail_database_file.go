package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetailDatabaseFile 数据库数据文件详情
type DetailDatabaseFile struct {
	Source *DatabaseFileSource `json:"source"`

	// 文件URL，用户私有数据中心为项目路径、公共数据场景为obs地址
	Url string `json:"url"`

	// 数据库文件所在项目id，仅文件为数据中心时填写
	EihealthProjectId *string `json:"eihealth_project_id,omitempty"`

	// 数据库文件写入状态
	Status *string `json:"status,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o DetailDatabaseFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailDatabaseFile struct{}"
	}

	return strings.Join([]string{"DetailDatabaseFile", string(data)}, " ")
}
