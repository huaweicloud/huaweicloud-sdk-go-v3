package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObjectMappingRequest Request Object
type ShowObjectMappingRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *QueryUserSelectedObjectInfoReq `json:"body,omitempty"`
}

func (o ShowObjectMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectMappingRequest struct{}"
	}

	return strings.Join([]string{"ShowObjectMappingRequest", string(data)}, " ")
}
