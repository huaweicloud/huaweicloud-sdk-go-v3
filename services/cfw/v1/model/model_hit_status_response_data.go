package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HitStatusResponseData struct {

	// RuleHitStatusListVO
	Records *[]RuleHitStatusListVo `json:"records,omitempty"`
}

func (o HitStatusResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HitStatusResponseData struct{}"
	}

	return strings.Join([]string{"HitStatusResponseData", string(data)}, " ")
}
