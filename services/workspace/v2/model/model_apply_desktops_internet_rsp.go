package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyDesktopsInternetRsp 创建云办公带宽响应。
type ApplyDesktopsInternetRsp struct {

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 【CBC回调】云运营平台CBC获取到JobId后，会使用getJobEndpoint当做URL，调用云服务，查询获取Job结果。
	GetJobEndpoint *string `json:"getJobEndpoint,omitempty"`

	// 【CBC回调】在线开通最大时间，在maxProvisionTime时间范围内，CBC会周期性的查询云服务开通结果；超过maxProvisionTime还没有开通成功，CBC会发失败工单，人工去分析处理。
	MaxProvisionTime *int32 `json:"maxProvisionTime,omitempty"`

	// 【CBC回调】开通最小时间（云服务最快开通时长，或一般开通时长）。获取到JobId后，经过minProvisionTime时间后，才来查询获取云服务开通结果。如果为空，云运营平台获取到JobId后，就去查询云服务开通结果。
	MinProvisionTime *int32 `json:"minProvisionTime,omitempty"`

	// 【CBC回调】云运营平台会使用getJobEndpoint(Job查询接口)、每隔periodicQueryTime时间去查询云服务开通结果。
	PeriodicQueryTime *int32 `json:"periodicQueryTime,omitempty"`
}

func (o ApplyDesktopsInternetRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDesktopsInternetRsp struct{}"
	}

	return strings.Join([]string{"ApplyDesktopsInternetRsp", string(data)}, " ")
}
