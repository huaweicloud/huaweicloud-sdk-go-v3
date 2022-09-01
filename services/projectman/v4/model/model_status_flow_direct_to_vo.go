package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type StatusFlowDirectToVo struct {

	//  父状态的名称
	ParentName *string `json:"parent_name,omitempty" xml:"parent_name"`

	// 父状态的类型
	ParentType *string `json:"parent_type,omitempty" xml:"parent_type"`

	// 状态id
	StatusId *string `json:"status_id,omitempty" xml:"status_id"`

	// 状态名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 状态类型
	StatusType *string `json:"status_type,omitempty" xml:"status_type"`

	// 是否已开启状态流转， true： 开启, false 没开启
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// 父状态的uuid
	ParentId *string `json:"parent_id,omitempty" xml:"parent_id"`
}

func (o StatusFlowDirectToVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusFlowDirectToVo struct{}"
	}

	return strings.Join([]string{"StatusFlowDirectToVo", string(data)}, " ")
}
