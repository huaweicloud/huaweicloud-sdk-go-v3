package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryOrgVmrResultDto 查询到的云会议室列表。
type QueryOrgVmrResultDto struct {

	// 云会议室的ID。 > 对应[[创建会议](https://support.huaweicloud.com/api-meeting/meeting_21_0014.html)](tag:hws)[[创建会议](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0014.html)](tag:hk)接口中的vmrID。
	Id *string `json:"id,omitempty"`

	// 云会议室的固定会议ID。 > 对应[[创建会议](https://support.huaweicloud.com/api-meeting/meeting_21_0014.html)](tag:hws)[[创建会议](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0014.html)](tag:hk)接口中的vmrConferenceID。
	VmrId *string `json:"vmrId,omitempty"`

	// 云会议室名称。
	VmrName *string `json:"vmrName,omitempty"`

	// 云会议室套餐名称。
	VmrPkgName *string `json:"vmrPkgName,omitempty"`

	// 云会议室套餐会议并发方数。
	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty"`

	// 最大观众与会方数（仅网络研讨会有效）。
	MaxAudienceParties *int32 `json:"maxAudienceParties,omitempty"`

	Member *IdMarkDto `json:"member,omitempty"`

	Device *IdMarkDto `json:"device,omitempty"`

	// 云会议室状态。 * 0：正常 * 1：冻结 * 2：未分配
	Status *int32 `json:"status,omitempty"`

	// 到期时间的时间戳，单位毫秒。
	ExpireDate *int64 `json:"expireDate,omitempty"`

	// 按次资源转商后，商用规格最大观众与会方数（仅网络研讨会有效）。
	CommercialMaxAudienceParties *int32 `json:"commercialMaxAudienceParties,omitempty"`
}

func (o QueryOrgVmrResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryOrgVmrResultDto struct{}"
	}

	return strings.Join([]string{"QueryOrgVmrResultDto", string(data)}, " ")
}
