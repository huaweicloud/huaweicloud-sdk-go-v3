package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性IP信息
type PublicipResult struct {
	Eip *EipResult `json:"eip,omitempty"`
}

func (o PublicipResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipResult struct{}"
	}

	return strings.Join([]string{"PublicipResult", string(data)}, " ")
}
