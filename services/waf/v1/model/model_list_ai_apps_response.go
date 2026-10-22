package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiAppsResponse Response Object
type ListAiAppsResponse struct {

	// 总条目数
	Total *int64 `json:"total,omitempty"`

	// AI应用列表
	Items          *[]AiAppItem `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAiAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiAppsResponse struct{}"
	}

	return strings.Join([]string{"ListAiAppsResponse", string(data)}, " ")
}
