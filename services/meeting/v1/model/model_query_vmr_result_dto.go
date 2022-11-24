package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户查询的云会议室或者个人会议ID详情。
type QueryVmrResultDto struct {

	// 云会议室的ID。 > 对应[[创建会议](https://support.huaweicloud.com/api-meeting/meeting_21_0014.html)](tag:hws)[[创建会议](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0014.html)](tag:hk)接口中的vmrID。
	Id *string `json:"id,omitempty"`

	// 云会议室的固定会议ID或者个人会议ID。 > 对应[[创建会议](https://support.huaweicloud.com/api-meeting/meeting_21_0014.html)](tag:hws)[[创建会议](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0014.html)](tag:hk)接口中的vmrConferenceID。
	VmrId *string `json:"vmrId,omitempty"`

	// 云会议室名称。
	VmrName *string `json:"vmrName,omitempty"`

	// 来宾密码。
	GustPwd *string `json:"gustPwd,omitempty"`

	// 来宾与会链接。
	GustJoinUrl *string `json:"gustJoinUrl,omitempty"`

	// 主持人密码。
	ChairPwd *string `json:"chairPwd,omitempty"`

	// 主持人与会链接。
	ChairJoinUrl *string `json:"chairJoinUrl,omitempty"`

	// 允许来宾先入会。
	AllowGustFirst *bool `json:"allowGustFirst,omitempty"`

	// 云会议室被使用后是否通知会议室所有者。
	GustFirstNotice *bool `json:"gustFirstNotice,omitempty"`

	// VMR模式。 * 0: 个人会议ID * 1: 云会议室 * 2: 网络研讨会
	VmrMode *int32 `json:"vmrMode,omitempty"`

	// 云会议室套餐包的id，仅云会议室返回。
	VmrPkgId *string `json:"vmrPkgId,omitempty"`

	// 云会议室套餐包的名称，仅云会议室返回。
	VmrPkgName *string `json:"vmrPkgName,omitempty"`

	// 云会议室套餐包的会议并发方数，仅云会议室返回。
	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty"`

	// 云会议室状态。 * 0：正常 * 1：停用 * 2：未分配
	Status *int32 `json:"status,omitempty"`
}

func (o QueryVmrResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVmrResultDto struct{}"
	}

	return strings.Join([]string{"QueryVmrResultDto", string(data)}, " ")
}
