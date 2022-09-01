package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCloudConnectionResponse struct {
	CloudConnection *CloudConnection `json:"cloud_connection,omitempty" xml:"cloud_connection"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudConnectionResponse", string(data)}, " ")
}
