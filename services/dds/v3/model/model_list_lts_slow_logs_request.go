package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsSlowLogsRequest Request Object
type ListLtsSlowLogsRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	Body *ListLtsSlowLogsRequestBody `json:"body,omitempty"`
}

func (o ListLtsSlowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowLogsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsSlowLogsRequest", string(data)}, " ")
}
