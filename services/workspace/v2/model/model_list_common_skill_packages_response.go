package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonSkillPackagesResponse Response Object
type ListCommonSkillPackagesResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 技能包列表项。
	Items *[]SkillPackageItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommonSkillPackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonSkillPackagesResponse struct{}"
	}

	return strings.Join([]string{"ListCommonSkillPackagesResponse", string(data)}, " ")
}
