package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApplicationEndpointAttributesResponse struct {
	// 请求的唯一标识ID。

	RequestId *string `json:"request_id,omitempty"`

	Attributes     *ListApplicationEndpointAttributesResponseBodyAttributes `json:"attributes,omitempty"`
	HttpStatusCode int                                                      `json:"-"`
}

func (o ListApplicationEndpointAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointAttributesResponse", string(data)}, " ")
}
