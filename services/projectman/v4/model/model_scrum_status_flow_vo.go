package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScrumStatusFlowVo struct {

	// 流转数据的uuid
	Id *string `json:"id,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`

	// 状态id
	StatusId *int32 `json:"status_id,omitempty"`

	// 流转到的数据
	DirectTo *[]ScrumStatusFlowDirectToVo `json:"direct_to,omitempty"`
}

func (o ScrumStatusFlowVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScrumStatusFlowVo struct{}"
	}

	return strings.Join([]string{"ScrumStatusFlowVo", string(data)}, " ")
}
