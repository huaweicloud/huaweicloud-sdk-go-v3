package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemovePackageRegionReq 移除技能包区域请求。
type RemovePackageRegionReq struct {

	// 要移除的区域标识列表。
	Regions []string `json:"regions"`
}

func (o RemovePackageRegionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemovePackageRegionReq struct{}"
	}

	return strings.Join([]string{"RemovePackageRegionReq", string(data)}, " ")
}
