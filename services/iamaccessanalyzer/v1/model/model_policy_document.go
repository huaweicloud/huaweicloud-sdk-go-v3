package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyDocument 该策略JSON格式策略文档。
type PolicyDocument struct {
}

func (o PolicyDocument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDocument struct{}"
	}

	return strings.Join([]string{"PolicyDocument", string(data)}, " ")
}
