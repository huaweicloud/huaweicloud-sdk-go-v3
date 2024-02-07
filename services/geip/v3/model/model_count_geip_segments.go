package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountGeipSegments struct {

	// Geip Segment Count
	Count int32 `json:"count"`
}

func (o CountGeipSegments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGeipSegments struct{}"
	}

	return strings.Join([]string{"CountGeipSegments", string(data)}, " ")
}
