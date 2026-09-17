package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObsBucketAclRequest Request Object
type ShowObsBucketAclRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// OBS桶名称
	BucketName string `json:"bucket_name"`
}

func (o ShowObsBucketAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObsBucketAclRequest struct{}"
	}

	return strings.Join([]string{"ShowObsBucketAclRequest", string(data)}, " ")
}
