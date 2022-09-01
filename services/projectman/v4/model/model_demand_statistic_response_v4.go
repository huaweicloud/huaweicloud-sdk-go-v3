package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 需求概览信息
type DemandStatisticResponseV4 struct {

	// 已关闭数量
	ClosedNum *int32 `json:"closed_num,omitempty" xml:"closed_num"`

	// 模块
	Module *string `json:"module,omitempty" xml:"module"`

	// 新建的数量
	NewNum *int32 `json:"new_num,omitempty" xml:"new_num"`

	// 开发中的数量
	ProcessNum *int32 `json:"process_num,omitempty" xml:"process_num"`

	// 已拒绝数量
	RejectedNum *int32 `json:"rejected_num,omitempty" xml:"rejected_num"`

	// 已解决数量
	SolvedNum *int32 `json:"solved_num,omitempty" xml:"solved_num"`

	// 测试中的数量
	TestNum *int32 `json:"test_num,omitempty" xml:"test_num"`

	// 总数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o DemandStatisticResponseV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemandStatisticResponseV4 struct{}"
	}

	return strings.Join([]string{"DemandStatisticResponseV4", string(data)}, " ")
}
