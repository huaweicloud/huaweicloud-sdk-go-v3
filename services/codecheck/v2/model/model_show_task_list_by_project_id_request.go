package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTaskListByProjectIdRequest struct {
	// 分页索引，偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的数量,每页最多显示100条

	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowTaskListByProjectIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskListByProjectIdRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskListByProjectIdRequest", string(data)}, " ")
}
