package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateRouteTableResponse struct {
	Association *Association `json:"association,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	XClientToken   *string `json:"X-Client-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"AssociateRouteTableResponse", string(data)}, " ")
}
