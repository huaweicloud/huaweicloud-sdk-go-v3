package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmssStep struct {

	// 测试步骤
	ExpectResult *string `json:"expectResult,omitempty"`

	// 预期结果对象
	TestStep *string `json:"testStep,omitempty"`
}

func (o TmssStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmssStep struct{}"
	}

	return strings.Join([]string{"TmssStep", string(data)}, " ")
}
