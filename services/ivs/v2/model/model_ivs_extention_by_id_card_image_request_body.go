package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IvsExtentionByIdCardImageRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsExtentionByIdCardImageRequestBodyData `json:"data"`
}

func (o IvsExtentionByIdCardImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsExtentionByIdCardImageRequestBody struct{}"
	}

	return strings.Join([]string{"IvsExtentionByIdCardImageRequestBody", string(data)}, " ")
}
