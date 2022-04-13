package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConcatInfo struct {
	// 拼接任务输入源地址。

	Inputs *[]ObsObjInfo `json:"inputs,omitempty"`
}

func (o ConcatInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConcatInfo struct{}"
	}

	return strings.Join([]string{"ConcatInfo", string(data)}, " ")
}
