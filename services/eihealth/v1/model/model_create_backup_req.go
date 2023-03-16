package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据归档请求参数
type CreateBackupReq struct {

	// 归档描述，最大长度为1000
	Description *string `json:"description,omitempty"`

	// 归档名称,最大长度为100
	Name string `json:"name"`

	// 归档路径集
	SubPaths []string `json:"sub_paths"`

	StorageType *StorageType `json:"storage_type,omitempty"`

	// 是否删除已归档数据
	DeleteArchivedData *bool `json:"delete_archived_data,omitempty"`
}

func (o CreateBackupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupReq struct{}"
	}

	return strings.Join([]string{"CreateBackupReq", string(data)}, " ")
}
