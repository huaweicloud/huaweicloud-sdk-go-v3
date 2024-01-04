package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AzFlavors 支持规格列表
type AzFlavors struct {

	// 可用区code
	AzCode *string `json:"az_code,omitempty"`

	// 可用区名称
	AzName *string `json:"az_name,omitempty"`

	// master节点支持的规格列表
	Master *[]Flavor `json:"master,omitempty"`

	// core节点支持的规格列表
	Core *[]Flavor `json:"core,omitempty"`

	// task节点支持的规格列表
	Task *[]Flavor `json:"task,omitempty"`
}

func (o AzFlavors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzFlavors struct{}"
	}

	return strings.Join([]string{"AzFlavors", string(data)}, " ")
}
