package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateNatTagResponse Response Object
type DeletePrivateNatTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePrivateNatTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatTagResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatTagResponse", string(data)}, " ")
}
