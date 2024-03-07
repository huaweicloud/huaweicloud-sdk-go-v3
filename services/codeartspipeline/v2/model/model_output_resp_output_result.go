package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputRespOutputResult struct {

	// output名称
	Key *string `json:"key,omitempty"`

	// output值
	Value *string `json:"value,omitempty"`
}

func (o OutputRespOutputResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputRespOutputResult struct{}"
	}

	return strings.Join([]string{"OutputRespOutputResult", string(data)}, " ")
}
