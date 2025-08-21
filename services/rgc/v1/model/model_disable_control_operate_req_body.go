package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableControlOperateReqBody 操作控制策略的请求体。
type DisableControlOperateReqBody struct {

	// 控制策略ID。
	Identifier string `json:"identifier"`

	// 组织单元的ID信息。
	TargetIdentifier string `json:"target_identifier"`
}

func (o DisableControlOperateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableControlOperateReqBody struct{}"
	}

	return strings.Join([]string{"DisableControlOperateReqBody", string(data)}, " ")
}
