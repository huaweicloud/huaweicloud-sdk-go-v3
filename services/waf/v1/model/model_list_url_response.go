package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUrlResponse Response Object
type ListUrlResponse struct {

	// 所有请求的每秒请求数
	Qps            *int32 `json:"qps,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUrlResponse struct{}"
	}

	return strings.Join([]string{"ListUrlResponse", string(data)}, " ")
}
