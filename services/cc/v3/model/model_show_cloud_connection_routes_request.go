package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCloudConnectionRoutesRequest struct {
	// 云连接路由实例ID。

	Id string `json:"id"`
}

func (o ShowCloudConnectionRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudConnectionRoutesRequest struct{}"
	}

	return strings.Join([]string{"ShowCloudConnectionRoutesRequest", string(data)}, " ")
}
