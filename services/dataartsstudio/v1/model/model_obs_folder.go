package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsFolder struct {

	// 目录名称
	FolderName *string `json:"folder_name,omitempty"`

	// 目录的guid
	FolderGuid *string `json:"folder_guid,omitempty"`

	// 目录的唯一标识名称
	FolderQualifiedName *string `json:"folder_qualified_name,omitempty"`

	// 对象总数
	ObjectCount *int32 `json:"object_count,omitempty"`

	// 数据量
	DataSize *float64 `json:"data_size,omitempty"`
}

func (o ObsFolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsFolder struct{}"
	}

	return strings.Join([]string{"ObsFolder", string(data)}, " ")
}
