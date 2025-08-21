package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableControlOperateReqBody 操作控制策略的请求体。
type EnableControlOperateReqBody struct {

	// 控制策略ID。
	Identifier string `json:"identifier"`

	// 组织单元的ID信息。
	TargetIdentifier string `json:"target_identifier"`

	// 策略参数。
	Parameters *[]EnableControlParameters `json:"parameters,omitempty"`
}

func (o EnableControlOperateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableControlOperateReqBody struct{}"
	}

	return strings.Join([]string{"EnableControlOperateReqBody", string(data)}, " ")
}
