package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Recommendation struct {

	// 标签名称。
	Name *string `json:"name,omitempty"`
}

func (o Recommendation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Recommendation struct{}"
	}

	return strings.Join([]string{"Recommendation", string(data)}, " ")
}
