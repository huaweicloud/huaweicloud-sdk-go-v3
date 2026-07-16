package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nfs nfs方式的挂载卷。
type Nfs struct {

	// nfs服务端路径，如：“10.10.10.10:/example/path”。
	NfsServerPath *string `json:"nfs_server_path,omitempty"`

	// 挂载到训练容器中的路径，如：“/example/path”。
	LocalPath *string `json:"local_path,omitempty"`

	// nfs挂载卷在容器中是否只读。
	ReadOnly *bool `json:"read_only,omitempty"`
}

func (o Nfs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nfs struct{}"
	}

	return strings.Join([]string{"Nfs", string(data)}, " ")
}
