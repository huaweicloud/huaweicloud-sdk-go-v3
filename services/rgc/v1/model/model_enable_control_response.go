package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableControlResponse Response Object
type EnableControlResponse struct {

	// 控制策略的操作ID。
	ControlOperateRequestId *string `json:"control_operate_request_id,omitempty"`
	HttpStatusCode          int     `json:"-"`
}

func (o EnableControlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableControlResponse struct{}"
	}

	return strings.Join([]string{"EnableControlResponse", string(data)}, " ")
}
