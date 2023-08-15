package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessConfigResponse Response Object
type ListAccessConfigResponse struct {

	// 日志接入列表
	Result *[]AccessConfigInfo `json:"result,omitempty"`

	// 日志接入总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAccessConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessConfigResponse struct{}"
	}

	return strings.Join([]string{"ListAccessConfigResponse", string(data)}, " ")
}
