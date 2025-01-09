package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopResponse Response Object
type ResizeDesktopResponse struct {

	// 错误码，失败时返回。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 创建云桌面总任务ID，CBC根据此ID定期查询任务是否成功
	CbcJobId *string `json:"cbcJobId,omitempty"`

	// 云运营平台CBC获取到JobId后，会使用getJobEndpoint当做URL，调用云服务，查询获取Job结果
	GetJobEndpoint *string `json:"getJobEndpoint,omitempty"`

	// 在线开通最大时间 在maxProvisionTime时间范围内，CBC会周期性的查询云服务开通结果；超过maxProvisionTime还没有开通成功，CBC会发失败工单，人工去分析处理。 单位：分钟。 如果为空，CBC默认为6小时。 取值范围（0,43200]，即30天。
	MaxProvisionTime *int32 `json:"maxProvisionTime,omitempty"`

	// 开通最小时间（云服务最快开通时长，或一般开通时长） 获取到JobId后，经过minProvisionTime时间后，才来查询获取云服务开通结果。如果为空，云运营平台获取到JobId后，就去查询云服务开通结果。 单位：分钟。 取值范围：(0, 43200)
	MinProvisionTime *int32 `json:"minProvisionTime,omitempty"`

	// Job周期性查询时间，默认2分钟查询一次 云运营平台会使用getJobEndpoint(Job查询接口)、每隔periodicQueryTime时间去查询云服务开通结果。 单位：分钟。 如果为空，则使用CBC默认的间隔时间（1分钟，2分钟，4分钟……15分钟）来查询云服务开通结果。 取值范围：(0, 43200)
	PeriodicQueryTime *int32 `json:"periodicQueryTime,omitempty"`

	// 变更订单错误处理策略。cbc调用返回值。设置为 NO_WORKORDER。云运营平台会认为无法开通成功，退费给客户后，不会再发运维工单给云服务，而由云服务自己去闭环处理对应问题。
	ErrorPolicy *string `json:"error_policy,omitempty"`

	// 按需桌面变更规格返回的任务信息（jobs字段后续会下线，请使用job_id字段）。
	Jobs *[]ResizeDesktopJobResponse `json:"jobs,omitempty"`

	// 变更规格任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeDesktopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopResponse struct{}"
	}

	return strings.Join([]string{"ResizeDesktopResponse", string(data)}, " ")
}
