package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicipResult 弹性IP信息
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
