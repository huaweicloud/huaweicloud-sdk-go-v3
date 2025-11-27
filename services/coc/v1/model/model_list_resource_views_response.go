package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceViewsResponse Response Object
type ListResourceViewsResponse struct {

	// 视图聚合的资源数据。
	Data           *[]ListViewResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListResourceViewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceViewsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceViewsResponse", string(data)}, " ")
}
