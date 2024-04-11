package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerResponse Response Object
type ChangeCloudPhoneServerResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 服务器id。
	ServerId *string `json:"server_id,omitempty"`

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeCloudPhoneServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerResponse struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerResponse", string(data)}, " ")
}
