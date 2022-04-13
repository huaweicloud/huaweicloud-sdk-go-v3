package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgencyAssumedby struct {
	User *AgencyAssumedbyUser `json:"user"`
}

func (o AgencyAssumedby) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyAssumedby struct{}"
	}

	return strings.Join([]string{"AgencyAssumedby", string(data)}, " ")
}
