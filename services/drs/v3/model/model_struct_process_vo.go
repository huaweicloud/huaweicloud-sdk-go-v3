package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 灾备初始化进度信息体
type StructProcessVo struct {

	// 对象类型
	Type string `json:"type" xml:"type"`

	// 状态
	Status int32 `json:"status" xml:"status"`

	// 源对象数量
	SrcCount int32 `json:"src_count" xml:"src_count"`

	// 目标对象数量
	DstCount int32 `json:"dst_count" xml:"dst_count"`

	// 开始时间
	StartTime int64 `json:"start_time" xml:"start_time"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time"`
}

func (o StructProcessVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructProcessVo struct{}"
	}

	return strings.Join([]string{"StructProcessVo", string(data)}, " ")
}
