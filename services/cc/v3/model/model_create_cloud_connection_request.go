package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudConnectionRequest Request Object
type CreateCloudConnectionRequest struct {
	Body *CreateCloudConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudConnectionRequest", string(data)}, " ")
}
