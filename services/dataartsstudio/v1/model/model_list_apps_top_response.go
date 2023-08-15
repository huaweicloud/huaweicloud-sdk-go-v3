package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsTopResponse Response Object
type ListAppsTopResponse struct {

	// 统计信息详情列表
	Statistics     *[]StatisticForDetail `json:"statistics,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAppsTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsTopResponse struct{}"
	}

	return strings.Join([]string{"ListAppsTopResponse", string(data)}, " ")
}
