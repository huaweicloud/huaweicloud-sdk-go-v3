package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableControlResponse Response Object
type DisableControlResponse struct {

	// 控制策略的操作ID。
	ControlOperateRequestId *string `json:"control_operate_request_id,omitempty"`
	HttpStatusCode          int     `json:"-"`
}

func (o DisableControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableControlResponse struct{}"
	}

	return strings.Join([]string{"DisableControlResponse", string(data)}, " ")
}
