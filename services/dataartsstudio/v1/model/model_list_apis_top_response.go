package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisTopResponse Response Object
type ListApisTopResponse struct {

	// 统计信息详情列表
	Statistics     *[]StatisticForDetail `json:"statistics,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListApisTopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisTopResponse struct{}"
	}

	return strings.Join([]string{"ListApisTopResponse", string(data)}, " ")
}
