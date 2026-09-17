package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsObjectsRequest Request Object
type ListObsObjectsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// OBS桶名
	BucketName string `json:"bucket_name"`

	// 最大对象数量
	MaxKeys int32 `json:"max_keys"`

	// 起始对象名称
	Marker string `json:"marker"`

	// 对象前缀
	Prefix string `json:"prefix"`
}

func (o ListObsObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListObsObjectsRequest", string(data)}, " ")
}
