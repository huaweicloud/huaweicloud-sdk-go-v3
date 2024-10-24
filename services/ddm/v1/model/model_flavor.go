package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Flavor struct {

	// 规格id。
	Id string `json:"id"`

	// 资源规格编码。
	SpecCode string `json:"spec_code"`

	// CPU大小。例如：1表示1U。
	Vcpus string `json:"vcpus"`

	// 内存大小，单位:GB。
	Ram string `json:"ram"`

	// 可用区信息  normal：在售。 unsupported：暂不支持该规格。 sellout：售罄。
	AzInfos []AvailableZone `json:"az_infos"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
