package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetCloudPhoneResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id"`

	// 任务信息
	Jobs           []interface{} `json:"jobs"`
	HttpStatusCode int           `json:"-"`
}

func (o ResetCloudPhoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetCloudPhoneResponse struct{}"
	}

	return strings.Join([]string{"ResetCloudPhoneResponse", string(data)}, " ")
}
