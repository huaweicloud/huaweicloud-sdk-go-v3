package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReasonDto struct {

	// 审批通过/驳回原因
	Reason string `json:"reason"`
}

func (o ReasonDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReasonDto struct{}"
	}

	return strings.Join([]string{"ReasonDto", string(data)}, " ")
}
