package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataServiceInstanceAccesslogsResponse Response Object
type ListDataServiceInstanceAccesslogsResponse struct {

	// 访问日志数量。
	Number *int32 `json:"number,omitempty"`

	// 访问日志列表。
	Records        *[]InstanceAccesslog `json:"records,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListDataServiceInstanceAccesslogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataServiceInstanceAccesslogsResponse struct{}"
	}

	return strings.Join([]string{"ListDataServiceInstanceAccesslogsResponse", string(data)}, " ")
}
