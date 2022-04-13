package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {
	// 实例规格信息。

	Flavors        *[]Flavor `json:"flavors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
