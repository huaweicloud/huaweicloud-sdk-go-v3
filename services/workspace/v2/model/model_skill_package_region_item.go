package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkillPackageRegionItem 技能包区域详情项（用于技能包详情中的 regions 列表）。
type SkillPackageRegionItem struct {

	// 区域标识。
	Region *string `json:"region,omitempty"`

	// OBS 桶名。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// OBS 对象键。
	ObsObjectKey *string `json:"obs_object_key,omitempty"`

	UploadStatus *UploadStatusEnum `json:"upload_status,omitempty"`
}

func (o SkillPackageRegionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkillPackageRegionItem struct{}"
	}

	return strings.Join([]string{"SkillPackageRegionItem", string(data)}, " ")
}
