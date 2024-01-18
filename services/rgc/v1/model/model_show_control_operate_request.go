package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowControlOperateRequest Request Object
type ShowControlOperateRequest struct {

	// 操作控制策略的请求ID。
	OperationControlStatusId string `json:"operation_control_status_id"`
}

func (o ShowControlOperateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowControlOperateRequest struct{}"
	}

	return strings.Join([]string{"ShowControlOperateRequest", string(data)}, " ")
}
