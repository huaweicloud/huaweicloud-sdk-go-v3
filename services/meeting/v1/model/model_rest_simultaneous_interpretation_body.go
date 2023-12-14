package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSimultaneousInterpretationBody 开启/关闭同声传译请求体
type RestSimultaneousInterpretationBody struct {

	// * 0：停止同声传译 * 1：启动同声传译
	SimultaneousInterpretation int32 `json:"simultaneousInterpretation"`
}

func (o RestSimultaneousInterpretationBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSimultaneousInterpretationBody struct{}"
	}

	return strings.Join([]string{"RestSimultaneousInterpretationBody", string(data)}, " ")
}
