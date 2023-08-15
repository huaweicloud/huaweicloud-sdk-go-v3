package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTransitIpTagResponse Response Object
type DeleteTransitIpTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTransitIpTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransitIpTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteTransitIpTagResponse", string(data)}, " ")
}
