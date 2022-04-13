package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteCloudConnectionRequest struct {
	// 云连接实例ID。

	Id string `json:"id"`
}

func (o DeleteCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudConnectionRequest", string(data)}, " ")
}
