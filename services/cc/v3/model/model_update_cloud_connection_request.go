package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCloudConnectionRequest struct {

	// 云连接实例ID。
	Id string `json:"id" xml:"id"`

	Body *UpdateCloudConnectionRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateCloudConnectionRequest", string(data)}, " ")
}
