package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssistAuthApplyObjectsRequest 更新辅助认证应用对象列表。
type AssistAuthApplyObjectsRequest struct {

	// 需要新增的应用对象。
	Add *[]ApplyObjects `json:"add,omitempty"`

	// 要移除的应用对象。
	Delete *[]ApplyObjects `json:"delete,omitempty"`
}

func (o AssistAuthApplyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssistAuthApplyObjectsRequest struct{}"
	}

	return strings.Join([]string{"AssistAuthApplyObjectsRequest", string(data)}, " ")
}
