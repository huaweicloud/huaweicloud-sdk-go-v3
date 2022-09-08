package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询云手机服务器猎豹
type ListCloudPhoneServersResponseBody struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id"`

	// 云手机服务器信息
	Servers []interface{} `json:"servers"`
}

func (o ListCloudPhoneServersResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServersResponseBody struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServersResponseBody", string(data)}, " ")
}
