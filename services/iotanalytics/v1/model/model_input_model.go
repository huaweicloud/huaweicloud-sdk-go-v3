package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InputModel struct {

	// 参数名称，正则：\"^[A-Za-z][A-Za-z_]{0,31}$\"
	Name string `json:"name" xml:"name"`

	PropertyReference *PropertyReferenceModel `json:"property_reference" xml:"property_reference"`
}

func (o InputModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputModel struct{}"
	}

	return strings.Join([]string{"InputModel", string(data)}, " ")
}
