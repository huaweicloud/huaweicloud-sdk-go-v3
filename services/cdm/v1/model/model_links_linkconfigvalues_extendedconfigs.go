package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinksLinkconfigvaluesExtendedconfigs 扩展配置，请参见extended-configs参数说明。
type LinksLinkconfigvaluesExtendedconfigs struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o LinksLinkconfigvaluesExtendedconfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksLinkconfigvaluesExtendedconfigs struct{}"
	}

	return strings.Join([]string{"LinksLinkconfigvaluesExtendedconfigs", string(data)}, " ")
}
