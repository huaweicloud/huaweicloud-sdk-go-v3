package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmHistoriesRequest Request Object
type ListAlarmHistoriesRequest struct {

	// 告警ID,以al开头，后跟22位由字母或数字组成的字符串
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警规则名称
	Name *string `json:"name,omitempty"`

	// 告警规则状态, ok为正常，alarm为告警，invalid为已失效
	Status *string `json:"status,omitempty"`

	// 告警级别, 1为紧急，2为重要，3为次要，4为提示
	Level *int32 `json:"level,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace *string `json:"namespace,omitempty"`

	// 告警资源ID，多维度情况按字母升序排列并使用逗号分隔
	ResourceId *string `json:"resource_id,omitempty"`

	// 查询告警记录的起始时间，例如：2022-02-10T10:05:46+08:00
	From *string `json:"from,omitempty"`

	// 查询告警记录的截止时间，例如：2022-02-10T10:05:47+08:00
	To *string `json:"to,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}
