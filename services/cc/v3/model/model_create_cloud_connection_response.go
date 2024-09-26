package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudConnectionResponse Response Object
type CreateCloudConnectionResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CloudConnection *CloudConnection `json:"cloud_connection"`
	HttpStatusCode  int              `json:"-"`
}

func (o CreateCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudConnectionResponse", string(data)}, " ")
}
