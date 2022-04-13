package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流计算
type DtStreamRequest struct {
	// 输入参数，最多支持10个

	Inputs []InputRequest `json:"inputs"`
}

func (o DtStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DtStreamRequest struct{}"
	}

	return strings.Join([]string{"DtStreamRequest", string(data)}, " ")
}
