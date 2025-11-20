package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenApiSpecSpec **参数解释**： 集群访问的地址 **约束限制**： 不涉及
type OpenApiSpecSpec struct {
	Eip *EipSpec `json:"eip,omitempty"`

	// **参数解释**： 是否动态创建 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	IsDynamic *bool `json:"IsDynamic,omitempty"`
}

func (o OpenApiSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApiSpecSpec struct{}"
	}

	return strings.Join([]string{"OpenApiSpecSpec", string(data)}, " ")
}
