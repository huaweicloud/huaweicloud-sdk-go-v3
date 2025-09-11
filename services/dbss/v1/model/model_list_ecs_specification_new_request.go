package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcsSpecificationNewRequest Request Object
type ListEcsSpecificationNewRequest struct {

	// **参数解释**： 镜像类型。 **约束限制**： 以取值范围为准 **取值范围**： - 1：运维 - 2：加密 - 3：审计 **默认取值**：   不传则默认为审计
	ImageBusinessType *int32 `json:"image_business_type,omitempty"`
}

func (o ListEcsSpecificationNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcsSpecificationNewRequest struct{}"
	}

	return strings.Join([]string{"ListEcsSpecificationNewRequest", string(data)}, " ")
}
