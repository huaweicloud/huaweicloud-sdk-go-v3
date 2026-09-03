package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonSkillsResponse Response Object
type ListCommonSkillsResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 公共技能列表项。
	Items *[]SkillListItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommonSkillsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonSkillsResponse struct{}"
	}

	return strings.Join([]string{"ListCommonSkillsResponse", string(data)}, " ")
}
