package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新云连接实例的请求体。
type UpdateCloudConnectionRequestBody struct {
	CloudConnection *UpdateCloudConnection `json:"cloud_connection" xml:"cloud_connection"`
}

func (o UpdateCloudConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCloudConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCloudConnectionRequestBody", string(data)}, " ")
}
