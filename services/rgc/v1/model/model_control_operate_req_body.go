package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ControlOperateReqBody 操作control请求体。
type ControlOperateReqBody struct {

	// 控制策略ID。
	Identifier string `json:"identifier"`

	// 组织单元的ID信息。
	TargetIdentifier string `json:"target_identifier"`
}

func (o ControlOperateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlOperateReqBody struct{}"
	}

	return strings.Join([]string{"ControlOperateReqBody", string(data)}, " ")
}
