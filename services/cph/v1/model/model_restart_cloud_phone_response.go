package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartCloudPhoneResponse Response Object
type RestartCloudPhoneResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务信息
	Jobs           *[]PhoneJob `json:"jobs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o RestartCloudPhoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartCloudPhoneResponse struct{}"
	}

	return strings.Join([]string{"RestartCloudPhoneResponse", string(data)}, " ")
}
