package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTasksReq This is a auto create Body Object
type DeleteTasksReq struct {

	// 待删除的任务ID列表
	Ids []string `json:"ids"`
}

func (o DeleteTasksReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTasksReq struct{}"
	}

	return strings.Join([]string{"DeleteTasksReq", string(data)}, " ")
}
