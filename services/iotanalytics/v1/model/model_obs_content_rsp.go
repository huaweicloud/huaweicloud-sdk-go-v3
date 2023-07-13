package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsContentRsp OBS数据源配置内容
type ObsContentRsp struct {

	// 桶名称
	BucketName *string `json:"bucketName,omitempty"`

	// 租户的AK
	Ak *string `json:"ak,omitempty"`

	// 租户的SK
	Sk *string `json:"sk,omitempty"`
}

func (o ObsContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsContentRsp struct{}"
	}

	return strings.Join([]string{"ObsContentRsp", string(data)}, " ")
}
