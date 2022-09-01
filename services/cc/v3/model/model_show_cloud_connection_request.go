package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCloudConnectionRequest struct {

	// 云连接实例ID。
	Id string `json:"id" xml:"id"`
}

func (o ShowCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowCloudConnectionRequest", string(data)}, " ")
}
