package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSkillPackageResp 查询技能包详情响应。
type ShowSkillPackageResp struct {

	// 技能包id。
	Id *string `json:"id,omitempty"`

	// 所属技能id。
	SkillId *string `json:"skill_id,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`

	// 版本修订号。
	Revision *int32 `json:"revision,omitempty"`

	// 技能包 SHA256 哈希值。
	PackageHash *string `json:"package_hash,omitempty"`

	// 技能包大小（字节）。
	PackageSize *int64 `json:"package_size,omitempty"`

	PackageStatus *PackageStatusEnum `json:"package_status,omitempty"`

	// 上传者。
	UploadedBy *string `json:"uploaded_by,omitempty"`

	// 上传者角色。
	UploadedRole *string `json:"uploaded_role,omitempty"`

	// 区域详情列表。
	Regions *[]SkillPackageRegionItem `json:"regions,omitempty"`

	// 备注。
	Remark *string `json:"remark,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ShowSkillPackageResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillPackageResp struct{}"
	}

	return strings.Join([]string{"ShowSkillPackageResp", string(data)}, " ")
}
