package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IvsStandardByNameAndIdRequestBody struct {
	Meta *Meta `json:"meta" xml:"meta"`

	Data *IvsStandardByNameAndIdRequestBodyData `json:"data" xml:"data"`
}

func (o IvsStandardByNameAndIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByNameAndIdRequestBody struct{}"
	}

	return strings.Join([]string{"IvsStandardByNameAndIdRequestBody", string(data)}, " ")
}
