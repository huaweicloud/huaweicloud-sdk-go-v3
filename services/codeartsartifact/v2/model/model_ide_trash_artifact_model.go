package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdeTrashArtifactModel 回收站信息
type IdeTrashArtifactModel struct {

	// **参数解释**: 仓库id，格式为{region}_{domainId}_{format}_{sequence}。可以从私有依赖库首页->仓库概览->仓库地址 url 中获取，最后两个\"/\"中间的字符串即为仓库id。 **约束限制**: 根据仓库id格式中region, domainId需要为有效值，format有效值为:npm|go|pypi|rpm|composer|maven|debian|conan|nuget|docker2|cocoapods|ohpm, sequence取值根据套餐不同有不同上限值。 **取值范围**: 不涉及。 **默认取值**: 无。
	Id string `json:"id"`

	// **参数解释**: 制品类型。 **约束限制**: 不涉及。 **取值范围**: maven2|docker|npm|go|pypi|rpm|composer|debian|conan|nuget|docker2|cocoapods|ohpm|generic。 **默认取值**: 无。
	Format string `json:"format"`

	// **参数解释**: 当前仓库状态。 **约束限制**: 不涉及。 **取值范围**: active：正常。 trash：废弃。 delete：删除。 **默认取值**: 无。
	Status string `json:"status"`

	// **参数解释**: 待还原的仓库路径。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	Uri string `json:"uri"`
}

func (o IdeTrashArtifactModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdeTrashArtifactModel struct{}"
	}

	return strings.Join([]string{"IdeTrashArtifactModel", string(data)}, " ")
}
