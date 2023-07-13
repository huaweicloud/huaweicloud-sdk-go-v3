package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StyleExtraMetaEditComponents struct {
	AdditionalProperties *StyleExtraMetaAdditionalProperties `json:"additionalProperties,omitempty"`
}

func (o StyleExtraMetaEditComponents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMetaEditComponents struct{}"
	}

	return strings.Join([]string{"StyleExtraMetaEditComponents", string(data)}, " ")
}
