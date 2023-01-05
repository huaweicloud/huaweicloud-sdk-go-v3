package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDbObjectCollectionStatusRequest struct {

	// 任务id
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 对象信息采集的id，指的是提交查询对象接口返回的id
	QueryId string `json:"query_id"`
}

func (o ShowDbObjectCollectionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectCollectionStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDbObjectCollectionStatusRequest", string(data)}, " ")
}
