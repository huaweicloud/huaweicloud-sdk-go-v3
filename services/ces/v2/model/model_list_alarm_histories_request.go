package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmHistoriesRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 告警ID,以al开头，后跟22位由字母或数字组成的字符串

	AlarmId *string `json:"alarm_id,omitempty"`
	// 告警规则名称

	Name *string `json:"name,omitempty"`
	// 告警规则状态

	Status *string `json:"status,omitempty"`
	// 告警规则等级

	Level *int32 `json:"level,omitempty"`
	// 服务的命名空间

	Namespace *string `json:"namespace,omitempty"`
	// 告警资源ID，多维度情况使用逗号分隔

	ResourceId *string `json:"resource_id,omitempty"`
	// 通过时间筛选traces的起始时间(不包括传入时间)，UTC时间

	From *string `json:"from,omitempty"`
	// 通过时间筛选traces的终止时间(不包括传入时间)，UTC时间

	To *string `json:"to,omitempty"`
	// 偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 希望的查询的数据量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}
