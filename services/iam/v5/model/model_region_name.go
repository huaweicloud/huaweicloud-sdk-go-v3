package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionName 区域名称。
type RegionName struct {
}

func (o RegionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionName struct{}"
	}

	return strings.Join([]string{"RegionName", string(data)}, " ")
}
