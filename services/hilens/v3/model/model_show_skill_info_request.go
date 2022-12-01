package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSkillInfoRequest struct {

	// 技能ID
	SkillId string `json:"skill_id"`

	// 技能审核状态状态，1表示审核通过，2表示审核不通过，0表示待审核
	Status *int32 `json:"status,omitempty"`

	// 技能版本
	Version *string `json:"version,omitempty"`
}

func (o ShowSkillInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSkillInfoRequest", string(data)}, " ")
}
