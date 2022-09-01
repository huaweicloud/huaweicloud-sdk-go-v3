package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CountDevicesResponse struct {
	ProductTemplates *ProductTemplatesCalculation `json:"product_templates,omitempty" xml:"product_templates"`

	Products *ProductsCalculation `json:"products,omitempty" xml:"products"`

	Devices        *DevicesCalculation `json:"devices,omitempty" xml:"devices"`
	HttpStatusCode int                 `json:"-"`
}

func (o CountDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountDevicesResponse struct{}"
	}

	return strings.Join([]string{"CountDevicesResponse", string(data)}, " ")
}
