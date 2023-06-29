package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublisherResponse Response Object
type ListPublisherResponse struct {

	// 返回值
	Result *[]PublisherVo `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublisherResponse struct{}"
	}

	return strings.Join([]string{"ListPublisherResponse", string(data)}, " ")
}
