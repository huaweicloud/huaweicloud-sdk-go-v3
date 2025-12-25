package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorParsersResponse Response Object
type ListCollectorParsersResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 相关描述信息
	Records        *[]Parser `json:"records,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCollectorParsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorParsersResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorParsersResponse", string(data)}, " ")
}
