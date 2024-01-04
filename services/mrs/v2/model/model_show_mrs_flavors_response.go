package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMrsFlavorsResponse Response Object
type ShowMrsFlavorsResponse struct {

	// 版本名称
	VersionName *string `json:"version_name,omitempty"`

	// 不同可用区支持的规格列表
	AvailableFlavors *[]AzFlavors `json:"available_flavors,omitempty"`
	HttpStatusCode   int          `json:"-"`
}

func (o ShowMrsFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMrsFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowMrsFlavorsResponse", string(data)}, " ")
}
