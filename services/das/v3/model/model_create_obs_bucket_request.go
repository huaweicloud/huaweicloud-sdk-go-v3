package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObsBucketRequest Request Object
type CreateObsBucketRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *CreateObsBucketRequestBody `json:"body,omitempty"`
}

func (o CreateObsBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObsBucketRequest struct{}"
	}

	return strings.Join([]string{"CreateObsBucketRequest", string(data)}, " ")
}
