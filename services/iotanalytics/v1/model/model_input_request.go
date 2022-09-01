package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InputRequest struct {

	// 参数名，正则： \"^[A-Za-z][A-Za-z_]{0,31}$\"
	Name string `json:"name" xml:"name"`

	PropertyReference *PropertyReferenceReq `json:"property_reference" xml:"property_reference"`
}

func (o InputRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InputRequest struct{}"
	}

	return strings.Join([]string{"InputRequest", string(data)}, " ")
}
