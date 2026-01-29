package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectRelatedRepositoryRequest Request Object
type ShowProjectRelatedRepositoryRequest struct {

	// **参数解释**： 项目名称，支持模糊查询。 **约束限制**： 不涉及。 **取值范围**： 最大200个字符。 **默认取值**： 不涉及。
	SearchName *string `json:"search_name,omitempty"`

	// **参数解释**： 分页查询的起始位置。 **约束限制**： 不涉及。 **取值范围**： 0-2147483647 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页查询每页的数据量。 **约束限制**： 不涉及。 **取值范围**： 1-2147483647 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 仓库ID，格式为{region}\\_{domainId}\\_{format}_{sequence}。可以从“私有依赖库首页 > 仓库概览 > 仓库地址url”中获取，最后两个“/”中间的字符串即为仓库ID。如仓库地址：https://devrepo.devcloud.abcde.abc.xyz.com/artgalaxy/abcde_09d2ca2f5080d5b60f51c00ae5bad0a0_maven_2_50/，其中abcde_09d2ca2f5080d5b60f51c00ae5bad0a0_maven_2_50即为仓库ID。 **约束限制**： 不涉及。 **取值范围**： 仓库ID格式中的format支持以下值： - npm - go - pypi - rpm - composer - maven - debian - conan - nuget - docker2 - cocoapods **默认取值**： 不涉及。
	RepoId *string `json:"repo_id,omitempty"`
}

func (o ShowProjectRelatedRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectRelatedRepositoryRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectRelatedRepositoryRequest", string(data)}, " ")
}
