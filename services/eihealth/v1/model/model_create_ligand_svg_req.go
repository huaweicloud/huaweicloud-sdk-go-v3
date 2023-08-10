package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLigandSvgReq 生成svg请求体
type CreateLigandSvgReq struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 尺寸
	Size *int32 `json:"size,omitempty"`

	// svg高度
	Height *int32 `json:"height,omitempty"`

	// svg宽度
	Width *int32 `json:"width,omitempty"`

	// alerts
	Alerts *string `json:"alerts,omitempty"`

	// ncols
	Ncols *int32 `json:"ncols,omitempty"`

	// bgopacity
	Bgopacity *float32 `json:"bgopacity,omitempty"`

	// bgcolor
	Bgcolor *string `json:"bgcolor,omitempty"`

	// fgcolor
	Fgcolor *string `json:"fgcolor,omitempty"`

	// ccolor
	Ccolor *string `json:"ccolor,omitempty"`

	// ncolor
	Ncolor *string `json:"ncolor,omitempty"`

	// ocolor
	Ocolor *string `json:"ocolor,omitempty"`
}

func (o CreateLigandSvgReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLigandSvgReq struct{}"
	}

	return strings.Join([]string{"CreateLigandSvgReq", string(data)}, " ")
}
