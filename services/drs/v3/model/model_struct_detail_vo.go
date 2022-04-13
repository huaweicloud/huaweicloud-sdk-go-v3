package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 灾备初始化对象详情信息体
type StructDetailVo struct {
	// 进度

	Progress *int32 `json:"progress,omitempty"`
	// 源数据库名称

	SrcDB *string `json:"src_DB,omitempty"`
	// 源对象名称

	SrcTB *string `json:"src_TB,omitempty"`
	// 目标数据库名称

	DstDB *string `json:"dst_DB,omitempty"`
	// 目标对象名称

	DstTB *string `json:"dst_TB,omitempty"`
}

func (o StructDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructDetailVo struct{}"
	}

	return strings.Join([]string{"StructDetailVo", string(data)}, " ")
}
