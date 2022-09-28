package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceRsp struct {
	Status *TaskInstanceStatusRsp `json:"status,omitempty"`

	Metadata *TaskInstanceMetadataRsp `json:"metadata,omitempty"`

	Spec *TaskInstanceSpecRsp `json:"spec,omitempty"`
}

func (o TaskInstanceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceRsp", string(data)}, " ")
}
