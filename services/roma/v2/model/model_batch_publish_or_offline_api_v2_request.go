package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPublishOrOfflineApiV2Request Request Object
type BatchPublishOrOfflineApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// - online：发布 - offline：下线
	Action string `json:"action"`

	Body *ApiBatchPublish `json:"body,omitempty"`
}

func (o BatchPublishOrOfflineApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPublishOrOfflineApiV2Request struct{}"
	}

	return strings.Join([]string{"BatchPublishOrOfflineApiV2Request", string(data)}, " ")
}
