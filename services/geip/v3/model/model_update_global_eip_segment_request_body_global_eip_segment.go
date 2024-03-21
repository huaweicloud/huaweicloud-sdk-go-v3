package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipSegmentRequestBodyGlobalEipSegment 更新全域弹性公网IP段请求对象
type UpdateGlobalEipSegmentRequestBodyGlobalEipSegment struct {

	// - 功能说明：全域弹性公网IP段名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// - 功能说明：用户自定义的资源描述 - 约束：   - 值的长度最大512字符，由数字、字母、中文、_(下划线)、-（中划线）、.（点）组成。
	Description *string `json:"description,omitempty"`
}

func (o UpdateGlobalEipSegmentRequestBodyGlobalEipSegment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipSegmentRequestBodyGlobalEipSegment struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipSegmentRequestBodyGlobalEipSegment", string(data)}, " ")
}
