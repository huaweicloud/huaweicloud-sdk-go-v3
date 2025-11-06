package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryAlarmRulesRequest Request Object
type ListFactoryAlarmRulesRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 作业/规则名称。
	Name *string `json:"name,omitempty"`

	// 通知类型：  - 0：运行成功。 - 1：运行异常/失败。 - 3：未完成。 - 4：资源繁忙。 - 5：基线任务异常。 - 6：基线预警。 - 7：基线破线。 - 8：基线加剧。 - 9：保障作业预警时间未完成。 - 10：保障作业承诺时间未完成。 - 11：保障作业失败。 - 12：周期未完成。 - 13：运行取消。 - 14：失败作业重跑成功。 - 15：作业改动。 默认查询所有状态。
	RemindType *int32 `json:"remind_type,omitempty"`

	// 钉钉机器人名称。
	DingName *string `json:"ding_name,omitempty"`

	// 分页列表的页数，默认值为0。取值范围大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页返回结果，指定每页最大记录数。范围[1,100]，默认为10。  默认值：10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFactoryAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryAlarmRulesRequest", string(data)}, " ")
}
