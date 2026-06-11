package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadHistorySessionAnalyaseRequest Request Object
type UploadHistorySessionAnalyaseRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *UploadHistorySessionAnalyaseBody `json:"body,omitempty"`
}

func (o UploadHistorySessionAnalyaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadHistorySessionAnalyaseRequest struct{}"
	}

	return strings.Join([]string{"UploadHistorySessionAnalyaseRequest", string(data)}, " ")
}
