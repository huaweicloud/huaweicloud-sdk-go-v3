package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// notebook存储信息
type NotebookStorage struct {

	// notebook存储路径
	Path string `json:"path"`

	// 挂载路径，由于目前暂不支持自定义挂载，暂不开放
	MountPath *string `json:"mount_path,omitempty"`
}

func (o NotebookStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookStorage struct{}"
	}

	return strings.Join([]string{"NotebookStorage", string(data)}, " ")
}
