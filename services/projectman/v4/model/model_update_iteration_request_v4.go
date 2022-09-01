package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateIterationRequestV4 struct {

	// 开始时间，年-月-日
	BeginTime string `json:"begin_time" xml:"begin_time"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 结束时间，年-月-日
	EndTime string `json:"end_time" xml:"end_time"`

	// 标题
	Name string `json:"name" xml:"name"`

	// 迭代的状态，0 未开始 <--> 1 进行中<--> 2 结束 <--> 1<-->0, 状态不能跨状态更改
	Status *string `json:"status,omitempty" xml:"status"`

	// 迭代结束时，工作项的处理（close 所有的工作项关闭，empty 没有关闭的工作项 放在block里面），status更新为2时需要填写over_type
	OverType *string `json:"over_type,omitempty" xml:"over_type"`
}

func (o UpdateIterationRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIterationRequestV4 struct{}"
	}

	return strings.Join([]string{"UpdateIterationRequestV4", string(data)}, " ")
}
