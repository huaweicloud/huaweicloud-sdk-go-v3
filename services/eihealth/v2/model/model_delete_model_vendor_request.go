package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelVendorRequest Request Object
type DeleteModelVendorRequest struct {

	// **参数解释**： 模型供应商ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	VendorId string `json:"vendor_id"`
}

func (o DeleteModelVendorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelVendorRequest struct{}"
	}

	return strings.Join([]string{"DeleteModelVendorRequest", string(data)}, " ")
}
