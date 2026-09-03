package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSkillPackageReq 更新技能包信息请求。
type UpdateSkillPackageReq struct {
	PackageStatus *PackageStatusEnum `json:"package_status,omitempty"`

	// 备注。
	Remark *string `json:"remark,omitempty"`
}

func (o UpdateSkillPackageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSkillPackageReq struct{}"
	}

	return strings.Join([]string{"UpdateSkillPackageReq", string(data)}, " ")
}
