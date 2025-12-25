package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryComponentConfigurationResponse Response Object
type ListHistoryComponentConfigurationResponse struct {

	// 列表总数
	Count *int64 `json:"count,omitempty"`

	// 节点参数
	Records        *[]ComponentConfigurationInfo `json:"records,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListHistoryComponentConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryComponentConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryComponentConfigurationResponse", string(data)}, " ")
}
