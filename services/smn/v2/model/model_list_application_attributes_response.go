package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApplicationAttributesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// Application的唯一标识ID。
	ApplicationId *string `json:"application_id,omitempty" xml:"application_id"`

	Attributes     *ListApplicationAttributesResponseBodyAttributes `json:"attributes,omitempty" xml:"attributes"`
	HttpStatusCode int                                              `json:"-"`
}

func (o ListApplicationAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationAttributesResponse", string(data)}, " ")
}
