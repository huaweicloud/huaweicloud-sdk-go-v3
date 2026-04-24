package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubscriptionDetailResponse Response Object
type ShowSubscriptionDetailResponse struct {

	// 任务id
	Id *string `json:"id,omitempty"`

	// 任务名称 约束：任务名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。 - 最小长度：4 - 最大长度：50
	Name *string `json:"name,omitempty"`

	// 内网ip
	Ip *string `json:"ip,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务状态，取值： CONFIGURATION：配置中 CREATING：创建中 CREATE_FAILED：创建失败 STARTJOBING：启动中 STARTJOB_FAILED：任务启动失败 SUBSCRIPTION_STARTED：正常 SUBSCRIPTION_FAILED：异常 DELETED：已删除 FROZEN：冻结状态 REBUILD_NODE_STARTED：订阅任务恢复中 REBUILD_NODE_FAILED：订阅任务恢复失败 NODE_UPGRADE_START：升级开始 NODE_UPGRADE_COMPLETE：升级完成 NODE_UPGRADE_FAILED：升级失败
	Status *string `json:"status,omitempty"`

	SubscriptionDataType *SubscriptionDataType `json:"subscription_data_type,omitempty"`

	SourceEndpoint *SubscriptionEndpointInfo `json:"source_endpoint,omitempty"`

	// 创建时间，以时间戳表示
	CreatedTime *string `json:"created_time,omitempty"`

	// 开始时间，以时间戳表示
	BeginTime *string `json:"begin_time,omitempty"`

	// 当前时间，以时间戳表示
	NowTime *string `json:"now_time,omitempty"`

	// 链路类型，当前仅支持“mysql”
	EngineType *string `json:"engine_type,omitempty"`

	ChargeInfo *ChargeInfoVo `json:"charge_info,omitempty"`

	// 描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubscriptionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionDetailResponse", string(data)}, " ")
}
