package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNetworkInstanceRequest struct {
	Body *CreateNetworkInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateNetworkInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNetworkInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateNetworkInstanceRequest", string(data)}, " ")
}
