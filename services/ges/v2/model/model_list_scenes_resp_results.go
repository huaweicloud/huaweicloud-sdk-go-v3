package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListScenesRespResults struct {

	// 场景名。
	Scene *string `json:"scene,omitempty"`

	// application名称。
	Name *string `json:"name,omitempty"`

	// 参数列表。
	Params *[]ListScenesRespParams `json:"params,omitempty"`

	// 场景下应用的描述。
	Description *string `json:"description,omitempty"`
}

func (o ListScenesRespResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScenesRespResults struct{}"
	}

	return strings.Join([]string{"ListScenesRespResults", string(data)}, " ")
}
