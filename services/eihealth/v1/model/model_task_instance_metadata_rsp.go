package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceMetadataRsp struct {

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o TaskInstanceMetadataRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceMetadataRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceMetadataRsp", string(data)}, " ")
}
