package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVendorAccountRequest Request Object
type UpdateVendorAccountRequest struct {

	// **参数解释：** 唯一标识id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id string `json:"id"`

	Body *VendorAccountUpdateRequest `json:"body,omitempty"`
}

func (o UpdateVendorAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVendorAccountRequest struct{}"
	}

	return strings.Join([]string{"UpdateVendorAccountRequest", string(data)}, " ")
}
