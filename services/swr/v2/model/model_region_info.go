package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegionInfo struct {

	// 区域ID
	RegionID string `json:"regionID"`
}

func (o RegionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionInfo struct{}"
	}

	return strings.Join([]string{"RegionInfo", string(data)}, " ")
}
