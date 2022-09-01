package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListOffSiteRestoreTimesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 所需查询的日期，为yyyy-mm-dd字符串格式，时区为UTC。
	Date *string `json:"date,omitempty" xml:"date"`
}

func (o ListOffSiteRestoreTimesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOffSiteRestoreTimesRequest struct{}"
	}

	return strings.Join([]string{"ListOffSiteRestoreTimesRequest", string(data)}, " ")
}
