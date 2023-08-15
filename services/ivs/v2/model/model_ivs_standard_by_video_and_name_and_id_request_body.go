package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IvsStandardByVideoAndNameAndIdRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsStandardByVideoAndNameAndIdRequestBodyData `json:"data"`
}

func (o IvsStandardByVideoAndNameAndIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByVideoAndNameAndIdRequestBody struct{}"
	}

	return strings.Join([]string{"IvsStandardByVideoAndNameAndIdRequestBody", string(data)}, " ")
}
