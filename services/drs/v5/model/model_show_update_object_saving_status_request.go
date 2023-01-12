package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowUpdateObjectSavingStatusRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 保存对象接口返回的ID。
	QueryId string `json:"query_id"`
}

func (o ShowUpdateObjectSavingStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateObjectSavingStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowUpdateObjectSavingStatusRequest", string(data)}, " ")
}
