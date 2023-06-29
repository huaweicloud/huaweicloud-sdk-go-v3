package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2XEdgeAppByEdgeAppIdResponse Response Object
type DeleteV2XEdgeAppByEdgeAppIdResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2XEdgeAppByEdgeAppIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2XEdgeAppByEdgeAppIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2XEdgeAppByEdgeAppIdResponse", string(data)}, " ")
}
