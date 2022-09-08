package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCloudConnectionRequest struct {

	// 云连接实例ID。
	Id string `json:"id"`

	Body *UpdateCloudConnectionRequestBody `json:"body,omitempty"`
}

func (o UpdateCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateCloudConnectionRequest", string(data)}, " ")
}
