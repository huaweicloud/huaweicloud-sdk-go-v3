package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 新增/更新节目信息请求
type ProgramRequestBase struct {

	// 节目名称
	ProgramName string `json:"programName"`
}

func (o ProgramRequestBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgramRequestBase struct{}"
	}

	return strings.Join([]string{"ProgramRequestBase", string(data)}, " ")
}
