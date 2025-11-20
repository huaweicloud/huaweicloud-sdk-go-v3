package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncReplicationMetadataRequest Request Object
type SyncReplicationMetadataRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o SyncReplicationMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncReplicationMetadataRequest struct{}"
	}

	return strings.Join([]string{"SyncReplicationMetadataRequest", string(data)}, " ")
}
