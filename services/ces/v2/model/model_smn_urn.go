package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmnUrn struct {
}

func (o SmnUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnUrn struct{}"
	}

	return strings.Join([]string{"SmnUrn", string(data)}, " ")
}
