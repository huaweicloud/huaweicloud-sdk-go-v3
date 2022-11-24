package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTasksRequest struct {

	// 是否将本次扫描升级为专业版规格（￥99.00/次）
	Upgrade *bool `json:"upgrade,omitempty"`

	Body *CreateTasksRequestBody `json:"body,omitempty"`
}

func (o CreateTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateTasksRequest", string(data)}, " ")
}
