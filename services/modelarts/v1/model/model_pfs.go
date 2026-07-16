package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Pfs 训练作业obsfs挂载卷信息。
type Pfs struct {

	// obsfs的地址。如：“/test-bucket/path”。
	PfsPath *string `json:"pfs_path,omitempty"`

	// 挂载到训练容器中的路径，如：“/example/path”。
	LocalPath *string `json:"local_path,omitempty"`
}

func (o Pfs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pfs struct{}"
	}

	return strings.Join([]string{"Pfs", string(data)}, " ")
}
