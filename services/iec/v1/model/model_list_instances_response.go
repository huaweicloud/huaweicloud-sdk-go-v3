package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesResponse Response Object
type ListInstancesResponse struct {

	// 异常站点。
	ErrSites *[]string `json:"err_sites,omitempty"`

	// 边缘实例列表的总数。
	Count *int32 `json:"count,omitempty"`

	// 边缘实例列表。
	Servers        *[]Instance `json:"servers,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
