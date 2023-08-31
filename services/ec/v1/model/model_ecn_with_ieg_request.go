package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EcnWithIegRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	// 区域ID
	RegionId string `json:"region_id"`
}

func (o EcnWithIegRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EcnWithIegRequest struct{}"
	}

	return strings.Join([]string{"EcnWithIegRequest", string(data)}, " ")
}
