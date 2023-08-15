package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StyleExtraMetaModellingAlgorithm struct {
	AdditionalProperties *StyleExtraMetaAdditionalProperties1 `json:"additionalProperties,omitempty"`
}

func (o StyleExtraMetaModellingAlgorithm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMetaModellingAlgorithm struct{}"
	}

	return strings.Join([]string{"StyleExtraMetaModellingAlgorithm", string(data)}, " ")
}
