package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudLogPlatformResponse Response Object
type ListCloudLogPlatformResponse struct {

	// 行管账户列表
	Result         *[]CloudLogManagerVo `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListCloudLogPlatformResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudLogPlatformResponse struct{}"
	}

	return strings.Join([]string{"ListCloudLogPlatformResponse", string(data)}, " ")
}
