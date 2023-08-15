package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExportCloudPhoneDataResponse Response Object
type BatchExportCloudPhoneDataResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务信息
	Jobs           *[]PhoneJob `json:"jobs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o BatchExportCloudPhoneDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportCloudPhoneDataResponse struct{}"
	}

	return strings.Join([]string{"BatchExportCloudPhoneDataResponse", string(data)}, " ")
}
