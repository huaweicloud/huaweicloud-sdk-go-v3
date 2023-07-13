package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinksLinkconfigvalues 连接参数配置，请参见link-config-values参数说明
type LinksLinkconfigvalues struct {

	// 连接配置参数数据结构，请参见configs参数说明。
	Configs []Configs `json:"configs"`

	ExtendedConfigs *LinksLinkconfigvaluesExtendedconfigs `json:"extended-configs,omitempty"`

	// 校验器
	Validators *[]string `json:"validators,omitempty"`
}

func (o LinksLinkconfigvalues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksLinkconfigvalues struct{}"
	}

	return strings.Join([]string{"LinksLinkconfigvalues", string(data)}, " ")
}
