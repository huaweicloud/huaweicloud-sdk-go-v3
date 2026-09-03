package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSkillsResponse Response Object
type ListSkillsResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 企业自研技能列表项。
	Items *[]SkillListItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSkillsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSkillsResponse struct{}"
	}

	return strings.Join([]string{"ListSkillsResponse", string(data)}, " ")
}
