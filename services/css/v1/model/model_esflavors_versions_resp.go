package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 引擎列表。
type EsflavorsVersionsResp struct {
	// 引擎版本，支持5.5.1、6.2.3、6.5.4、7.1.1、7.6.2、7.9.3。

	Version *string `json:"version,omitempty"`
	// 规格列表。

	Flavors *[]EsflavorsVersionsFlavorsResp `json:"flavors,omitempty"`
}

func (o EsflavorsVersionsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsflavorsVersionsResp struct{}"
	}

	return strings.Join([]string{"EsflavorsVersionsResp", string(data)}, " ")
}
