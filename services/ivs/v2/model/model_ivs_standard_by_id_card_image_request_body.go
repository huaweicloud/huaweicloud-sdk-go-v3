package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IvsStandardByIdCardImageRequestBody struct {
	Meta *Meta `json:"meta" xml:"meta"`

	Data *IvsStandardByIdCardImageRequestBodyData `json:"data" xml:"data"`
}

func (o IvsStandardByIdCardImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByIdCardImageRequestBody struct{}"
	}

	return strings.Join([]string{"IvsStandardByIdCardImageRequestBody", string(data)}, " ")
}
