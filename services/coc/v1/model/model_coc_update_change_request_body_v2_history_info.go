package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocUpdateChangeRequestBodyV2HistoryInfo 历史信息
type CocUpdateChangeRequestBodyV2HistoryInfo struct {

	// 操作类型
	Action *string `json:"action,omitempty"`
}

func (o CocUpdateChangeRequestBodyV2HistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocUpdateChangeRequestBodyV2HistoryInfo struct{}"
	}

	return strings.Join([]string{"CocUpdateChangeRequestBodyV2HistoryInfo", string(data)}, " ")
}
