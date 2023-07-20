package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowCompareData 对比结果。
type FlowCompareData struct {

	// 源数据库名称。
	SrcDb *string `json:"src_db,omitempty"`

	// 源对象名称。
	SrcTb *string `json:"src_tb,omitempty"`

	// 目标数据库名称。
	DstDb *string `json:"dst_db,omitempty"`

	// 目标对象名称。
	DstTb *string `json:"dst_tb,omitempty"`

	// 进度，1-100。
	Progress *int32 `json:"progress,omitempty"`
}

func (o FlowCompareData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowCompareData struct{}"
	}

	return strings.Join([]string{"FlowCompareData", string(data)}, " ")
}
