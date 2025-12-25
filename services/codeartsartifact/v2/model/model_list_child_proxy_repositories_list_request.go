package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListChildProxyRepositoriesListRequest Request Object
type ListChildProxyRepositoriesListRequest struct {

	// **参数解释**: 仓库id，格式为{region}_{domainId}_{format}_{sequence}。可以从私有依赖库首页->仓库概览->仓库地址 url 中获取，最后两个\"/\"中间的字符串即为仓库id。 **约束限制**: 不涉及。 **取值范围**: 根据仓库id格式中region, domainId需要为有效值，format有效值为:npm|go|pypi|rpm|composer|maven|debian|conan|nuget|docker2|cocoapods|ohpm|generic|helm|conda, sequence取值根据套餐不同有不同上限值。 **默认取值**: 不涉及。
	RepoId string `json:"repo_id"`

	// **参数解释**: type。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o ListChildProxyRepositoriesListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChildProxyRepositoriesListRequest struct{}"
	}

	return strings.Join([]string{"ListChildProxyRepositoriesListRequest", string(data)}, " ")
}
