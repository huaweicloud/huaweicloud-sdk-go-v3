package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceContainerStatusRsp struct {

	// 重启次数
	RestartCount *int32 `json:"restart_count,omitempty"`
}

func (o TaskInstanceContainerStatusRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceContainerStatusRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceContainerStatusRsp", string(data)}, " ")
}
