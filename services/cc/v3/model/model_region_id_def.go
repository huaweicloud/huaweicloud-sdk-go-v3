package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionIdDef RegionID。
type RegionIdDef struct {
}

func (o RegionIdDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionIdDef struct{}"
	}

	return strings.Join([]string{"RegionIdDef", string(data)}, " ")
}
