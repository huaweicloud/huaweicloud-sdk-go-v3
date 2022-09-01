package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCloudPhoneServersResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 云手机服务器信息
	Servers        []interface{} `json:"servers" xml:"servers"`
	HttpStatusCode int           `json:"-"`
}

func (o ListCloudPhoneServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServersResponse struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServersResponse", string(data)}, " ")
}
