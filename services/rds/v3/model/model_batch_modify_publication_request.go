package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyPublicationRequest Request Object
type BatchModifyPublicationRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *BatchModifyPublicationsRequestBody `json:"body,omitempty"`
}

func (o BatchModifyPublicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyPublicationRequest struct{}"
	}

	return strings.Join([]string{"BatchModifyPublicationRequest", string(data)}, " ")
}
