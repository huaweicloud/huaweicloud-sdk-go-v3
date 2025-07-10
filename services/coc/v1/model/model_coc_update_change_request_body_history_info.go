package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocUpdateChangeRequestBodyHistoryInfo 变更历史信息。
type CocUpdateChangeRequestBodyHistoryInfo struct {

	// -| 操作类型。 · change_start_change_success：变更开始 · change_end_change_success：变更结束 · change_set_change_result_success：变更成功添加变更结果 · change_result_failed：变更失败添加变更结果 · change_cancel_change_success：变更取消
	Action *string `json:"action,omitempty"`
}

func (o CocUpdateChangeRequestBodyHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocUpdateChangeRequestBodyHistoryInfo struct{}"
	}

	return strings.Join([]string{"CocUpdateChangeRequestBodyHistoryInfo", string(data)}, " ")
}
