package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStrategyApplyObjectsRequestBody 批量操作策略应用对象请求
type UpdateStrategyApplyObjectsRequestBody struct {

	// 操作类型：1-批量删除 2-批量新增
	OperateType int32 `json:"operate_type"`

	// 批量删除的应用对象ID列表（operate_type为1时必填）
	DeleteIds *[]string `json:"delete_ids,omitempty"`

	// 批量新增的应用对象列表（operate_type为2时必填）
	AddObjects *[]ApplyObjectInfo `json:"add_objects,omitempty"`
}

func (o UpdateStrategyApplyObjectsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStrategyApplyObjectsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateStrategyApplyObjectsRequestBody", string(data)}, " ")
}
