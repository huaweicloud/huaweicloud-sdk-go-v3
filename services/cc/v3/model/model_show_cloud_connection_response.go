package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCloudConnectionResponse struct {
	CloudConnection *CloudConnection `json:"cloud_connection,omitempty"`
	// 请求ID。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudConnectionResponse", string(data)}, " ")
}
