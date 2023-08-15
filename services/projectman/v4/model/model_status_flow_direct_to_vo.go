package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusFlowDirectToVo
type StatusFlowDirectToVo struct {

	//  父状态的名称
	ParentName *string `json:"parent_name,omitempty"`

	// 父状态的类型
	ParentType *string `json:"parent_type,omitempty"`

	// 状态id
	StatusId *string `json:"status_id,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`

	// 状态类型
	StatusType *string `json:"status_type,omitempty"`

	// 是否已开启状态流转， true： 开启, false 没开启
	Enabled *bool `json:"enabled,omitempty"`

	// 父状态的uuid
	ParentId *string `json:"parent_id,omitempty"`
}

func (o StatusFlowDirectToVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusFlowDirectToVo struct{}"
	}

	return strings.Join([]string{"StatusFlowDirectToVo", string(data)}, " ")
}
