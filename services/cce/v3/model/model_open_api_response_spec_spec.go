package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenApiResponseSpecSpec struct {
	Eip *EipSpec `json:"eip,omitempty" xml:"eip"`

	// 是否动态创建
	IsDynamic *bool `json:"IsDynamic,omitempty" xml:"IsDynamic"`
}

func (o OpenApiResponseSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApiResponseSpecSpec struct{}"
	}

	return strings.Join([]string{"OpenApiResponseSpecSpec", string(data)}, " ")
}
