package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWorkspacesRequest Request Object
type BatchDeleteWorkspacesRequest struct {

	// DataArts Studio实例ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteWorkspacesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkspacesRequest", string(data)}, " ")
}
