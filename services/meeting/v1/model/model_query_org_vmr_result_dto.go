package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组织查询的vmr列表，不越权显示vmr的来宾密码，主席密码等
type QueryOrgVmrResultDto struct {
	// 唯一标识。 说明：对应会议管理->创建会议接口中的vmrID。

	Id *string `json:"id,omitempty"`
	// 云会议室ID。 说明：对应会议管理->创建会议接口中当vmrIDType等于0（固定ID）时返回数据的conferenceID 。

	VmrId *string `json:"vmrId,omitempty"`
	// 云会议室名称。

	VmrName *string `json:"vmrName,omitempty"`
	// 云会议室套餐名称。

	VmrPkgName *string `json:"vmrPkgName,omitempty"`
	// 云会议室套餐会议并发方数。

	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty"`

	Member *IdMarkDto `json:"member,omitempty"`

	Device *IdMarkDto `json:"device,omitempty"`
	// 云会议室状态。 * 0：正常 * 1：冻结 * 2：未分配

	Status *int32 `json:"status,omitempty"`
}

func (o QueryOrgVmrResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryOrgVmrResultDto struct{}"
	}

	return strings.Join([]string{"QueryOrgVmrResultDto", string(data)}, " ")
}
