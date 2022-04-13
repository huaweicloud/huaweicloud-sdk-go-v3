package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {
	// 计算类型规格组。

	ComputeFlavorGroups *[]ComputeFlavorGroupsInfo `json:"computeFlavorGroups,omitempty"`
	HttpStatusCode      int                        `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
