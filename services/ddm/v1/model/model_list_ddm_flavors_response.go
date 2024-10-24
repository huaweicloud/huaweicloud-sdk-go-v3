package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDdmFlavorsResponse Response Object
type ListDdmFlavorsResponse struct {

	// 规格组。
	FlavorGroups   *[]FlavorGroupInfo `json:"flavor_groups,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListDdmFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDdmFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListDdmFlavorsResponse", string(data)}, " ")
}
