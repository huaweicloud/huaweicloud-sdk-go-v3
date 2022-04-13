package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApiVersionDetailV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// API版本的编号

	VersionId string `json:"version_id"`
}

func (o ListApiVersionDetailV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionDetailV2Request struct{}"
	}

	return strings.Join([]string{"ListApiVersionDetailV2Request", string(data)}, " ")
}
