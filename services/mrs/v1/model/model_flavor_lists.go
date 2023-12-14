package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorLists 支持规格列表
type FlavorLists struct {

	// master节点支持的规格列表
	Master *[]string `json:"master,omitempty"`

	// core节点支持的规格列表
	Core *[]string `json:"core,omitempty"`

	// task节点支持的规格列表
	Task *[]string `json:"task,omitempty"`
}

func (o FlavorLists) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorLists struct{}"
	}

	return strings.Join([]string{"FlavorLists", string(data)}, " ")
}
