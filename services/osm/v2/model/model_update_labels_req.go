package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLabelsReq struct {

	// 标签名称
	Name string `json:"name"`

	// 颜色id
	Color string `json:"color"`
}

func (o UpdateLabelsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLabelsReq struct{}"
	}

	return strings.Join([]string{"UpdateLabelsReq", string(data)}, " ")
}
