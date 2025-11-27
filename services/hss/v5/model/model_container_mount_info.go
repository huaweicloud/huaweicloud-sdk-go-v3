package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerMountInfo struct {

	// 主机路径
	HostPath *string `json:"host_path,omitempty"`

	// 挂载路径
	MountPath *string `json:"mount_path,omitempty"`

	// 挂载传播方式
	MountPropagation *string `json:"mount_propagation,omitempty"`

	// 是否只读
	ReadOnly *bool `json:"read_only,omitempty"`

	// 卷名称
	MountName *string `json:"mount_name,omitempty"`

	// 子路径
	SubPath *string `json:"sub_path,omitempty"`

	// 子路径表达式
	SubPathExpr *string `json:"sub_path_expr,omitempty"`
}

func (o ContainerMountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerMountInfo struct{}"
	}

	return strings.Join([]string{"ContainerMountInfo", string(data)}, " ")
}
