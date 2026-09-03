package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportExportObsObjectsRequest Request Object
type ImportExportObsObjectsRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 最大返回对象数
	MaxKeys *int32 `json:"max_keys,omitempty"`

	// 标记
	Marker *string `json:"marker,omitempty"`

	// 前缀
	Prefix *string `json:"prefix,omitempty"`
}

func (o ImportExportObsObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportExportObsObjectsRequest struct{}"
	}

	return strings.Join([]string{"ImportExportObsObjectsRequest", string(data)}, " ")
}
