package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteV2XEdgeByV2xEdgeIdResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2XEdgeByV2xEdgeIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2XEdgeByV2xEdgeIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2XEdgeByV2xEdgeIdResponse", string(data)}, " ")
}
