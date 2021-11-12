package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteApiByVersionIdV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API版本的编号

	VersionId string `json:"version_id"`
}

func (o DeleteApiByVersionIdV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiByVersionIdV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiByVersionIdV2Request", string(data)}, " ")
}
