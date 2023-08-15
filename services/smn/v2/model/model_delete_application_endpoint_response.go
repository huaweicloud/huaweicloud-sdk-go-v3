package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationEndpointResponse Response Object
type DeleteApplicationEndpointResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteApplicationEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationEndpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationEndpointResponse", string(data)}, " ")
}
