package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSkillsResponse Response Object
type ListInstanceSkillsResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 实例绑定的技能列表项。
	Items *[]InstanceSkillItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInstanceSkillsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSkillsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSkillsResponse", string(data)}, " ")
}
