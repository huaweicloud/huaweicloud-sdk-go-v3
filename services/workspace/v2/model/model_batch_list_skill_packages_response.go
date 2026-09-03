package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListSkillPackagesResponse Response Object
type BatchListSkillPackagesResponse struct {

	// 返回的技能包结果总数。
	Count *int32 `json:"count,omitempty"`

	// 技能包查询结果列表。
	Items *[]BatchListSkillPackageItem `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchListSkillPackagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListSkillPackagesResponse struct{}"
	}

	return strings.Join([]string{"BatchListSkillPackagesResponse", string(data)}, " ")
}
