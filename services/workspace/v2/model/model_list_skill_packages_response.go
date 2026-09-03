package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSkillPackagesResponse Response Object
type ListSkillPackagesResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 技能包列表项。
	Items *[]SkillPackageItemVo `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSkillPackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSkillPackagesResponse struct{}"
	}

	return strings.Join([]string{"ListSkillPackagesResponse", string(data)}, " ")
}
