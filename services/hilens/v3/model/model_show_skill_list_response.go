package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSkillListResponse Response Object
type ShowSkillListResponse struct {

	// 技能列表
	Data *[]SkillInfo `json:"data,omitempty"`

	// 总数量
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSkillListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillListResponse struct{}"
	}

	return strings.Join([]string{"ShowSkillListResponse", string(data)}, " ")
}
