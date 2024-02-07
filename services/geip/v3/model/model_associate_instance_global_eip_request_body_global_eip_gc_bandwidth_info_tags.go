package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTags struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTags struct{}"
	}

	return strings.Join([]string{"AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTags", string(data)}, " ")
}
