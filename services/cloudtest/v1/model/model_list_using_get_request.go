package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsingGetRequest Request Object
type ListUsingGetRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 看板名称，精确匹配
	Name *string `json:"name,omitempty"`

	// 页码
	PageNumber *int32 `json:"page_number,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListUsingGetRequest", string(data)}, " ")
}
