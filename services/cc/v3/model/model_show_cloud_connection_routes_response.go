package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudConnectionRoutesResponse Response Object
type ShowCloudConnectionRoutesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CloudConnectionRoute *CloudConnectionRoute `json:"cloud_connection_route"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowCloudConnectionRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudConnectionRoutesResponse struct{}"
	}

	return strings.Join([]string{"ShowCloudConnectionRoutesResponse", string(data)}, " ")
}
