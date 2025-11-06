package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentKillSessionHistoryResponse Response Object
type ShowIntelligentKillSessionHistoryResponse struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 限流信息列表
	History *[]IntelligentKillSessionHistory `json:"history,omitempty"`

	// 历史记录数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowIntelligentKillSessionHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentKillSessionHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowIntelligentKillSessionHistoryResponse", string(data)}, " ")
}
