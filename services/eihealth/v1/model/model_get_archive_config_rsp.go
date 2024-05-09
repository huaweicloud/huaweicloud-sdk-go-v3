package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetArchiveConfigRsp struct {

	// 华为云项目
	RegionId *string `json:"region_id,omitempty"`

	// 该区域是否是当前设置的归档区域
	CurrentRegion *bool `json:"current_region,omitempty"`
}

func (o GetArchiveConfigRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetArchiveConfigRsp struct{}"
	}

	return strings.Join([]string{"GetArchiveConfigRsp", string(data)}, " ")
}
