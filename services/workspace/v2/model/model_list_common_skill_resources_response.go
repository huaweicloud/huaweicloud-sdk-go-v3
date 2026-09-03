package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonSkillResourcesResponse Response Object
type ListCommonSkillResourcesResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 技能绑定的资源列表项。
	Items *[]SkillResourceItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommonSkillResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonSkillResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListCommonSkillResourcesResponse", string(data)}, " ")
}
