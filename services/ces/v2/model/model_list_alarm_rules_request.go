package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmRulesRequest struct {

	// 发送的实体的MIME类型。默认使用application/json; charset=UTF-8。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 告警规则ID
	AlarmId *string `json:"alarm_id,omitempty" xml:"alarm_id"`

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name *string `json:"name,omitempty" xml:"name"`

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 告警资源ID，多维度情况按字母升序排列并使用逗号分隔
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesRequest", string(data)}, " ")
}
