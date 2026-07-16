package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSearchAlgorithmsSearchAlgoList struct {

	// 超参搜索算法的名称。
	Name *string `json:"name,omitempty"`

	// 超参搜索算法的参数列表。
	Params *[]ListSearchAlgorithmsParams `json:"params,omitempty"`

	// 超参搜索算法的描述。
	Description *string `json:"description,omitempty"`
}

func (o ListSearchAlgorithmsSearchAlgoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchAlgorithmsSearchAlgoList struct{}"
	}

	return strings.Join([]string{"ListSearchAlgorithmsSearchAlgoList", string(data)}, " ")
}
