package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagCloudConnectionRequest Request Object
type UntagCloudConnectionRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *UntagCloudConnectionRequestBody `json:"body,omitempty"`
}

func (o UntagCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"UntagCloudConnectionRequest", string(data)}, " ")
}
