package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdmetWithCustomRequest ADMET with custom请求体
type AdmetWithCustomRequest struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 用户已开启的自定义属性集合
	CustomProps *[]CustomProp `json:"custom_props,omitempty"`
}

func (o AdmetWithCustomRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdmetWithCustomRequest struct{}"
	}

	return strings.Join([]string{"AdmetWithCustomRequest", string(data)}, " ")
}
