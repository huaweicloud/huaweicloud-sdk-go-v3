package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIsapComponentsResponse Response Object
type ListIsapComponentsResponse struct {

	// 组件数量
	Count *int64 `json:"count,omitempty"`

	// 组件
	Records        *[]ComponentDetailInfo `json:"records,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListIsapComponentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIsapComponentsResponse struct{}"
	}

	return strings.Join([]string{"ListIsapComponentsResponse", string(data)}, " ")
}
