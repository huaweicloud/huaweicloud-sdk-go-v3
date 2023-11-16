package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullStagePluginsRelationVoAddables struct {

	// 额外属性1
	AdditionalProp1 *bool `json:"additionalProp1,omitempty"`

	// 额外属性2
	AdditionalProp2 *bool `json:"additionalProp2,omitempty"`

	// 额外属性3
	AdditionalProp3 *bool `json:"additionalProp3,omitempty"`
}

func (o FullStagePluginsRelationVoAddables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagePluginsRelationVoAddables struct{}"
	}

	return strings.Join([]string{"FullStagePluginsRelationVoAddables", string(data)}, " ")
}
