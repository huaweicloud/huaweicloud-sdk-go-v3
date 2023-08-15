package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopCloudPhoneResponse Response Object
type StopCloudPhoneResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务信息
	Jobs           *[]PhoneJob `json:"jobs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o StopCloudPhoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCloudPhoneResponse struct{}"
	}

	return strings.Join([]string{"StopCloudPhoneResponse", string(data)}, " ")
}
