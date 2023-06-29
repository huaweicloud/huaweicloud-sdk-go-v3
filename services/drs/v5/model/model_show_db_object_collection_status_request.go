package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDbObjectCollectionStatusRequest Request Object
type ShowDbObjectCollectionStatusRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 对象信息采集的ID，提交查询数据库对象信息接口返回的ID。
	QueryId string `json:"query_id"`
}

func (o ShowDbObjectCollectionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbObjectCollectionStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDbObjectCollectionStatusRequest", string(data)}, " ")
}
