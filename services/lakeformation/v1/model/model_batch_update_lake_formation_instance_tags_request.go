package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdateLakeFormationInstanceTagsRequest struct {

	// LakeFormation实例ID
	InstanceId string `json:"instance_id"`

	Body *BatchUpdateTagsRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateLakeFormationInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateLakeFormationInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateLakeFormationInstanceTagsRequest", string(data)}, " ")
}
