package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApplicationEndpointsResponse struct {
	// 请求的唯一标识ID。

	RequestId *string `json:"request_id,omitempty"`
	// 是否有下一页标识。

	NextPageFlag *bool `json:"next_page_flag,omitempty"`
	// Application_endpoint结构体数。

	Endpoints      *[]ApplicationEndpoint `json:"endpoints,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListApplicationEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointsResponse", string(data)}, " ")
}
