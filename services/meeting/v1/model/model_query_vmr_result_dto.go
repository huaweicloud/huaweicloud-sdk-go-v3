package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户查询的vmr详情
type QueryVmrResultDto struct {
	// 唯一标识。 说明：对应会议管理->创建会议接口中的vmrID。

	Id *string `json:"id,omitempty"`
	// 云会议室ID。 说明：对应会议管理->创建会议接口中当vmrIDType等于0（固定ID）时返回数据的conferenceID 。

	VmrId *string `json:"vmrId,omitempty"`
	// 云会议室名称

	VmrName *string `json:"vmrName,omitempty"`
	// 来宾密码

	GustPwd *string `json:"gustPwd,omitempty"`
	// 来宾与会链接

	GustJoinUrl *string `json:"gustJoinUrl,omitempty"`
	// 主席密码

	ChairPwd *string `json:"chairPwd,omitempty"`
	// 主席与会链接

	ChairJoinUrl *string `json:"chairJoinUrl,omitempty"`
	// 允许来宾先入会

	AllowGustFirst *bool `json:"allowGustFirst,omitempty"`
	// 云会议室被使用后是否通知会议室所有者

	GustFirstNotice *bool `json:"gustFirstNotice,omitempty"`
	// VMR模式 * 0: 个人会议ID * 1: 云会议室 * 2: 网络研讨会

	VmrMode *int32 `json:"vmrMode,omitempty"`
	// 云会议室套餐包的id，仅专用云会议室返回

	VmrPkgId *string `json:"vmrPkgId,omitempty"`
	// 云会议室套餐包的名称，仅专用云会议室返回

	VmrPkgName *string `json:"vmrPkgName,omitempty"`
	// 云会议室套餐包的会议并发方数，仅专用云会议室返回

	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty"`
	// 云会议室状态: * 0、正常 * 1、停用 * 2、未分配

	Status *int32 `json:"status,omitempty"`
}

func (o QueryVmrResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVmrResultDto struct{}"
	}

	return strings.Join([]string{"QueryVmrResultDto", string(data)}, " ")
}
