package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StructProcessVo 灾备初始化进度信息体
type StructProcessVo struct {

	// 对象类型
	Type string `json:"type"`

	// 状态
	Status int32 `json:"status"`

	// 源对象数量
	SrcCount int32 `json:"src_count"`

	// 目标对象数量
	DstCount int32 `json:"dst_count"`

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o StructProcessVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructProcessVo struct{}"
	}

	return strings.Join([]string{"StructProcessVo", string(data)}, " ")
}
