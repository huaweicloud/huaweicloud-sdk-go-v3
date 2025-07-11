package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopVolumesResponse Response Object
type AddDesktopVolumesResponse struct {

	// 增加磁盘任务id。
	JobId *string `json:"job_id,omitempty"`

	// 云运营平台CBC获取到JobId后，会使用getJobEndpoint当做URL，调用云服务，查询获取Job结果。
	GetJobEndpoint *string `json:"getJobEndpoint,omitempty"`

	// 在线开通最大时间。
	MaxProvisionTime *int32 `json:"maxProvisionTime,omitempty"`

	// 开通最小时间（云服务最快开通时长，或一般开通时长）。
	MinProvisionTime *int32 `json:"minProvisionTime,omitempty"`

	// Job周期性查询时间，默认1分钟查询一次。
	PeriodicQueryTime *int32 `json:"periodicQueryTime,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o AddDesktopVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopVolumesResponse struct{}"
	}

	return strings.Join([]string{"AddDesktopVolumesResponse", string(data)}, " ")
}
