package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataspacesResponse Response Object
type ListDataspacesResponse struct {

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 数据空间列表
	Records        *[]Dataspace `json:"records,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDataspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataspacesResponse struct{}"
	}

	return strings.Join([]string{"ListDataspacesResponse", string(data)}, " ")
}
