package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询Vmr套餐包分配数量结果
type QueryVmrPkgResResultDto struct {

	// 云会议室套餐包id。
	VmrPkgId *string `json:"vmrPkgId,omitempty" xml:"vmrPkgId"`

	// 云会议室套餐包名称。
	VmrName *string `json:"vmrName,omitempty" xml:"vmrName"`

	// 云会议室套餐方数。
	VmrPkgParties *int32 `json:"vmrPkgParties,omitempty" xml:"vmrPkgParties"`

	// 该云会议室套餐分配的总数。
	VmrPkgCount *int32 `json:"vmrPkgCount,omitempty" xml:"vmrPkgCount"`

	// 该套餐对应的云会议室已分配数量。
	VmrPkgUsedCount *int32 `json:"vmrPkgUsedCount,omitempty" xml:"vmrPkgUsedCount"`
}

func (o QueryVmrPkgResResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVmrPkgResResultDto struct{}"
	}

	return strings.Join([]string{"QueryVmrPkgResResultDto", string(data)}, " ")
}
