package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateP2PSiteNetworkRequest Request Object
type CreateP2PSiteNetworkRequest struct {
	Body *CreateP2PSiteNetworkRequestBody `json:"body,omitempty"`
}

func (o CreateP2PSiteNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2PSiteNetworkRequest struct{}"
	}

	return strings.Join([]string{"CreateP2PSiteNetworkRequest", string(data)}, " ")
}
