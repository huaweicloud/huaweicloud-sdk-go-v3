package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserTagDto struct {
	MatchField *string `json:"match_field,omitempty"`

	MatchFieldIndex *string `json:"match_field_index,omitempty"`
}

func (o UserTagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserTagDto struct{}"
	}

	return strings.Join([]string{"UserTagDto", string(data)}, " ")
}
