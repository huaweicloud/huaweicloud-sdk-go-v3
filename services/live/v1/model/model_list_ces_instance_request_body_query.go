package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCesInstanceRequestBodyQuery struct {

	// 维度名称
	DimName string `json:"dim_name"`

	// 维度id
	Id *string `json:"id,omitempty"`
}

func (o ListCesInstanceRequestBodyQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceRequestBodyQuery struct{}"
	}

	return strings.Join([]string{"ListCesInstanceRequestBodyQuery", string(data)}, " ")
}
