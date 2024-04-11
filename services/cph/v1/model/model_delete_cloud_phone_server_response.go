package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudPhoneServerResponse Response Object
type DeleteCloudPhoneServerResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteCloudPhoneServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudPhoneServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteCloudPhoneServerResponse", string(data)}, " ")
}
