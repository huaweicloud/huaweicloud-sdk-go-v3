package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShortGlobalEipSegment struct {

	// 全域弹性公网IP段的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：全域弹性公网IP段名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`
}

func (o ShortGlobalEipSegment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShortGlobalEipSegment struct{}"
	}

	return strings.Join([]string{"ShortGlobalEipSegment", string(data)}, " ")
}
