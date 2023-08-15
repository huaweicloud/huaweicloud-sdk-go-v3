package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScrumStatusFlowDirectToVo struct {

	// 流转数据的uuid
	Id *string `json:"id,omitempty"`

	// 状态id
	StatusId *int32 `json:"status_id,omitempty"`

	// 状态名
	Name *string `json:"name,omitempty"`

	// 是否开启流转
	Enabled *bool `json:"enabled,omitempty"`
}

func (o ScrumStatusFlowDirectToVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScrumStatusFlowDirectToVo struct{}"
	}

	return strings.Join([]string{"ScrumStatusFlowDirectToVo", string(data)}, " ")
}
