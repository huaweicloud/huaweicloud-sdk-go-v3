package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNoticeConfigsResponse Response Object
type ListNoticeConfigsResponse struct {

	// 配置的告警通知总数量
	Total *int32 `json:"total,omitempty"`

	// 配置的告警通知
	Items          *[]AlertNoticeConfigResponse `json:"items,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListNoticeConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticeConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListNoticeConfigsResponse", string(data)}, " ")
}
