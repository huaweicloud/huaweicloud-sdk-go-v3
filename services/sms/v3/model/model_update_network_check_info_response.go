package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNetworkCheckInfoResponse Response Object
type UpdateNetworkCheckInfoResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNetworkCheckInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkCheckInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateNetworkCheckInfoResponse", string(data)}, " ")
}
