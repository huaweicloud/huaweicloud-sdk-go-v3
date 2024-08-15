package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceResponse Response Object
type ListResourceResponse struct {

	// 资源列表
	Data *[]BatchListResourceResponseData `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceResponse struct{}"
	}

	return strings.Join([]string{"ListResourceResponse", string(data)}, " ")
}
