package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditHlsInfo struct {
	// 切片间隔。

	Interval *int32 `json:"interval,omitempty"`
}

func (o EditHlsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditHlsInfo struct{}"
	}

	return strings.Join([]string{"EditHlsInfo", string(data)}, " ")
}
