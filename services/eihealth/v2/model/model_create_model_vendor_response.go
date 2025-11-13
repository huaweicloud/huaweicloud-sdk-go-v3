package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelVendorResponse Response Object
type CreateModelVendorResponse struct {

	// **参数解释**： 模型供应商ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateModelVendorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelVendorResponse struct{}"
	}

	return strings.Join([]string{"CreateModelVendorResponse", string(data)}, " ")
}
