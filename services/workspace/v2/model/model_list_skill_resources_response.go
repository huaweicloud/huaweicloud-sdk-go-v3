package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSkillResourcesResponse Response Object
type ListSkillResourcesResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 技能绑定的资源列表项。
	Items *[]SkillResourceItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSkillResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSkillResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListSkillResourcesResponse", string(data)}, " ")
}
