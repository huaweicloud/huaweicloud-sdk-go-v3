package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPermissionResult 查询权限返回结果
type CheckPermissionResult struct {

	// 对应输入策略的检查结果
	CheckResult bool `json:"check_result"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 行过滤条件列表。
	DataFilters *[]string `json:"data_filters,omitempty"`

	// 列掩码条件列表。
	DataMasks *[]DataMaskInfo `json:"data_masks,omitempty"`
}

func (o CheckPermissionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPermissionResult struct{}"
	}

	return strings.Join([]string{"CheckPermissionResult", string(data)}, " ")
}
