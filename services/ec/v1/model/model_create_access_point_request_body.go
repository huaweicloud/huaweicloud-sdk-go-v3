package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAccessPointRequestBody struct {

	// 区域ID
	RegionId string `json:"region_id"`

	// 带宽
	BandwidthSize int32 `json:"bandwidth_size"`
}

func (o CreateAccessPointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessPointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAccessPointRequestBody", string(data)}, " ")
}
