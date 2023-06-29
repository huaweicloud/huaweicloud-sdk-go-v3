package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DtTransformRequest 转换计算
type DtTransformRequest struct {

	// 输入参数，最多支持10个
	Inputs *[]InputRequest `json:"inputs,omitempty"`

	Outputs *[]OutputRequest `json:"outputs,omitempty"`
}

func (o DtTransformRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DtTransformRequest struct{}"
	}

	return strings.Join([]string{"DtTransformRequest", string(data)}, " ")
}
