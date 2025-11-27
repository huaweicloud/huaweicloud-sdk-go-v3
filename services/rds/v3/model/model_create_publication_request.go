package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicationRequest Request Object
type CreatePublicationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreatePublicationsRequestBody `json:"body,omitempty"`
}

func (o CreatePublicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicationRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicationRequest", string(data)}, " ")
}
