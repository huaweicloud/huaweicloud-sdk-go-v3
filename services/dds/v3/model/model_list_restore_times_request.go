package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreTimesRequest Request Object
type ListRestoreTimesRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	// 所需查询的日期，为yyyy-mm-dd字符串格式，时区为UTC。
	Date string `json:"date"`
}

func (o ListRestoreTimesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreTimesRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreTimesRequest", string(data)}, " ")
}
