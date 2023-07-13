package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportCloudPhoneDataResponse Response Object
type BatchImportCloudPhoneDataResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务信息
	Jobs           *[]PhoneJob `json:"jobs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o BatchImportCloudPhoneDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportCloudPhoneDataResponse struct{}"
	}

	return strings.Join([]string{"BatchImportCloudPhoneDataResponse", string(data)}, " ")
}
