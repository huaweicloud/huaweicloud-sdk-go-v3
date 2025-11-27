package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemoteDbRequest Request Object
type ListRemoteDbRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ListRemoteDbRequestBody `json:"body,omitempty"`
}

func (o ListRemoteDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemoteDbRequest struct{}"
	}

	return strings.Join([]string{"ListRemoteDbRequest", string(data)}, " ")
}
