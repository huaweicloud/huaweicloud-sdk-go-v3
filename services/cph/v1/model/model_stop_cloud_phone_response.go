package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopCloudPhoneResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 任务信息
	Jobs           []interface{} `json:"jobs" xml:"jobs"`
	HttpStatusCode int           `json:"-"`
}

func (o StopCloudPhoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCloudPhoneResponse struct{}"
	}

	return strings.Join([]string{"StopCloudPhoneResponse", string(data)}, " ")
}
