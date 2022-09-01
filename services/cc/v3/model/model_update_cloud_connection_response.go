package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateCloudConnectionResponse struct {
	CloudConnection *CloudConnection `json:"cloud_connection,omitempty" xml:"cloud_connection"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateCloudConnectionResponse", string(data)}, " ")
}
