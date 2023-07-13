package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCloudConnectionResponse Response Object
type UpdateCloudConnectionResponse struct {
	CloudConnection *CloudConnection `json:"cloud_connection,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateCloudConnectionResponse", string(data)}, " ")
}
