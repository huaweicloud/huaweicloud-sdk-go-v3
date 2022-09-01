package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmHistoriesRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 服务提供的资源分组功能，创建的资源分组ID，如：rg1603107497873DK4O2pXbn。
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 告警规则ID。如：al1603088932912v98rGl1al。
	AlarmId *string `json:"alarm_id,omitempty" xml:"alarm_id"`

	// 告警规则名称，如alarm-test01。
	AlarmName *string `json:"alarm_name,omitempty" xml:"alarm_name"`

	// 告警历史的状态，取值为ok，alarm，insufficient_data； ok为正常，alarm为告警，insufficient_data为数据不足。
	AlarmStatus *string `json:"alarm_status,omitempty" xml:"alarm_status"`

	// 告警历史的告警级别，值为1,2,3,4；1为紧急，2为重要，3为次要，4为提示。
	AlarmLevel *string `json:"alarm_level,omitempty" xml:"alarm_level"`

	// 告警资源对应的命名空间，如ECS服务的资源命名空间为：SYS.ECS；各服务命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 查询告警历史的起始时间，UNIX时间戳，单位毫秒，如：1602501480905；from，to如果不进行赋值，则默认to是当前时间，from是当前时间减7天的时间戳。
	From *string `json:"from,omitempty" xml:"from"`

	// 查询告警历史的截止时间，UNIX时间戳，单位毫秒。from必须小于等于to，如：1603106280905；from，to如果不进行赋值，则默认to是当前时间，from是当前时间减7天的时间戳。
	To *string `json:"to,omitempty" xml:"to"`

	// 分页起始值，类型为integer，默认值为0。
	Start *string `json:"start,omitempty" xml:"start"`

	// 单次查询的条数限制，取值范围(0,100]，默认值为100， 用于限制结果数据条数。
	Limit *string `json:"limit,omitempty" xml:"limit"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}
