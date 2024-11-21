package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ControlOperateReqBody 操作控制策略的请求体。
type ControlOperateReqBody struct {

	// 控制策略ID。
	Identifier string `json:"identifier"`

	// 组织单元的ID信息。
	TargetIdentifier string `json:"target_identifier"`

	// 策略参数。
	Parameters *[]EnableControlParameters `json:"parameters,omitempty"`
}

func (o ControlOperateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlOperateReqBody struct{}"
	}

	return strings.Join([]string{"ControlOperateReqBody", string(data)}, " ")
}
