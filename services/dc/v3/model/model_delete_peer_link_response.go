package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePeerLinkResponse Response Object
type DeletePeerLinkResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePeerLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePeerLinkResponse struct{}"
	}

	return strings.Join([]string{"DeletePeerLinkResponse", string(data)}, " ")
}
