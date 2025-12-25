package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudLogManagesResponse Response Object
type ListCloudLogManagesResponse struct {

	// 行管账户列表
	Result         *[]CloudLogManagerVo `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListCloudLogManagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudLogManagesResponse struct{}"
	}

	return strings.Join([]string{"ListCloudLogManagesResponse", string(data)}, " ")
}
