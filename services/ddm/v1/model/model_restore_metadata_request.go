package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreMetadataRequest Request Object
type RestoreMetadataRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	Body *RestoreMetaData2ExistReq `json:"body,omitempty"`
}

func (o RestoreMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreMetadataRequest struct{}"
	}

	return strings.Join([]string{"RestoreMetadataRequest", string(data)}, " ")
}
