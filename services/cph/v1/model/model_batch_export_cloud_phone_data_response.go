package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchExportCloudPhoneDataResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 任务信息
	Jobs           []interface{} `json:"jobs" xml:"jobs"`
	HttpStatusCode int           `json:"-"`
}

func (o BatchExportCloudPhoneDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportCloudPhoneDataResponse struct{}"
	}

	return strings.Join([]string{"BatchExportCloudPhoneDataResponse", string(data)}, " ")
}
