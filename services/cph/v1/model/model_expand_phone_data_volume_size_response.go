package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandPhoneDataVolumeSizeResponse Response Object
type ExpandPhoneDataVolumeSizeResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务信息。
	Jobs           *[]PhoneJob `json:"jobs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ExpandPhoneDataVolumeSizeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandPhoneDataVolumeSizeResponse struct{}"
	}

	return strings.Join([]string{"ExpandPhoneDataVolumeSizeResponse", string(data)}, " ")
}
