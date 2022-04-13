package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IvsStandardByNameAndIdRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsStandardByNameAndIdRequestBodyData `json:"data"`
}

func (o IvsStandardByNameAndIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByNameAndIdRequestBody struct{}"
	}

	return strings.Join([]string{"IvsStandardByNameAndIdRequestBody", string(data)}, " ")
}
