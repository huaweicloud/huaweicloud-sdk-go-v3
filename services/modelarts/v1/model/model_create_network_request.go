package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNetworkRequest Request Object
type CreateNetworkRequest struct {
	Body *NetworkCreationRequest `json:"body,omitempty"`
}

func (o CreateNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNetworkRequest struct{}"
	}

	return strings.Join([]string{"CreateNetworkRequest", string(data)}, " ")
}
