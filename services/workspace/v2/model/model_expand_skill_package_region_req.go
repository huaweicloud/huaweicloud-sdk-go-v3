package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandSkillPackageRegionReq 扩展技能包区域请求（企业自研技能，含上传状态）。
type ExpandSkillPackageRegionReq struct {

	// 新 region 信息（含上传状态）。
	Regions []PackageRegionWithStatusInfo `json:"regions"`
}

func (o ExpandSkillPackageRegionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandSkillPackageRegionReq struct{}"
	}

	return strings.Join([]string{"ExpandSkillPackageRegionReq", string(data)}, " ")
}
