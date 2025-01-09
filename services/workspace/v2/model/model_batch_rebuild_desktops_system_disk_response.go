package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebuildDesktopsSystemDiskResponse Response Object
type BatchRebuildDesktopsSystemDiskResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 加密后的详细拒绝原因，用户可以自行调用STS服务的decode-authorization-message接口进行解密。
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`

	// 重建系统盘总任务id。
	JobId *string `json:"job_id,omitempty"`

	// 云运营平台CBC获取到JobId后，会使用getJobEndpoint当做URL，调用云服务，查询获取Job结果。
	GetJobEndpoint *string `json:"get_job_endpoint,omitempty"`

	// 在线开通最大时间。
	MaxProvisionTime *int32 `json:"max_provision_time,omitempty"`

	// 开通最小时间（云服务最快开通时长，或一般开通时长）。
	MinProvisionTime *int32 `json:"min_provision_time,omitempty"`

	// Job周期性查询时间。
	PeriodicQueryTime *int32 `json:"periodic_query_time,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o BatchRebuildDesktopsSystemDiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebuildDesktopsSystemDiskResponse struct{}"
	}

	return strings.Join([]string{"BatchRebuildDesktopsSystemDiskResponse", string(data)}, " ")
}
