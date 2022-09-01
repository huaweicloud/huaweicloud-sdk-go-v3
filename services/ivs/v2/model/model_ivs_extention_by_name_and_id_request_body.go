package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IvsExtentionByNameAndIdRequestBody struct {
	Meta *Meta `json:"meta" xml:"meta"`

	Data *IvsExtentionByNameAndIdRequestBodyData `json:"data" xml:"data"`
}

func (o IvsExtentionByNameAndIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsExtentionByNameAndIdRequestBody struct{}"
	}

	return strings.Join([]string{"IvsExtentionByNameAndIdRequestBody", string(data)}, " ")
}
