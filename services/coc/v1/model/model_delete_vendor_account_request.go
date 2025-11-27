package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVendorAccountRequest Request Object
type DeleteVendorAccountRequest struct {

	// **参数解释：** 唯一标识id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id string `json:"id"`
}

func (o DeleteVendorAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVendorAccountRequest struct{}"
	}

	return strings.Join([]string{"DeleteVendorAccountRequest", string(data)}, " ")
}
