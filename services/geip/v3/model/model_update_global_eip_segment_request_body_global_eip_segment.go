package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipSegmentRequestBodyGlobalEipSegment 更新全域弹性公网IP段请求对象
type UpdateGlobalEipSegmentRequestBodyGlobalEipSegment struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateGlobalEipSegmentRequestBodyGlobalEipSegment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipSegmentRequestBodyGlobalEipSegment struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipSegmentRequestBodyGlobalEipSegment", string(data)}, " ")
}
