package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCloudConnectionRequest Request Object
type UpdateCloudConnectionRequest struct {

	// 实例ID。
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
