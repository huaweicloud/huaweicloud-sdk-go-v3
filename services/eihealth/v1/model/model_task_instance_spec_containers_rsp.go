package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceSpecContainersRsp struct {
	Resources *TaskInstanceSpecConResourceRsp `json:"resources,omitempty"`
}

func (o TaskInstanceSpecContainersRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceSpecContainersRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceSpecContainersRsp", string(data)}, " ")
}
