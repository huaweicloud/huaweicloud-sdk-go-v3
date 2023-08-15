package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudConnectionResponse Response Object
type CreateCloudConnectionResponse struct {
	CloudConnection *CloudConnection `json:"cloud_connection,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudConnectionResponse", string(data)}, " ")
}
