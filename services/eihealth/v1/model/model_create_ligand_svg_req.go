package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLigandSvgReq 生成svg请求体
type CreateLigandSvgReq struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 分子骨架smiles表达式
	Scaffold *string `json:"scaffold,omitempty"`

	// 尺寸
	Size *int32 `json:"size,omitempty"`

	// svg高度
	Height *int32 `json:"height,omitempty"`

	// svg宽度
	Width *int32 `json:"width,omitempty"`

	// 高亮子结构编号
	Alerts *string `json:"alerts,omitempty"`

	// 显示的列数
	Ncols *int32 `json:"ncols,omitempty"`

	// 背景透明度
	Bgopacity *float32 `json:"bgopacity,omitempty"`

	// 背景颜色
	Bgcolor *string `json:"bgcolor,omitempty"`

	// 前景色
	Fgcolor *string `json:"fgcolor,omitempty"`

	// 碳颜色
	Ccolor *string `json:"ccolor,omitempty"`

	// 氮颜色
	Ncolor *string `json:"ncolor,omitempty"`

	// 氧颜色
	Ocolor *string `json:"ocolor,omitempty"`
}

func (o CreateLigandSvgReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLigandSvgReq struct{}"
	}

	return strings.Join([]string{"CreateLigandSvgReq", string(data)}, " ")
}
