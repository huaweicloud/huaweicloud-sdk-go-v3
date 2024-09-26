package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudConnectionResponse Response Object
type ShowCloudConnectionResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CloudConnection *CloudConnection `json:"cloud_connection"`
	HttpStatusCode  int              `json:"-"`
}

func (o ShowCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudConnectionResponse", string(data)}, " ")
}
