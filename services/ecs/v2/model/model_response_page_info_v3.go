package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponsePageInfoV3 struct {
	NextMarker *string `json:"next_marker,omitempty"`

	PreviousMarker *string `json:"previous_marker,omitempty"`
}

func (o ResponsePageInfoV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponsePageInfoV3 struct{}"
	}

	return strings.Join([]string{"ResponsePageInfoV3", string(data)}, " ")
}
