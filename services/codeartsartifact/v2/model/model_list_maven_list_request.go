package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMavenListRequest Request Object
type ListMavenListRequest struct {

	// **参数解释**: 项目ID，可以从调用API处获取，也可以从控制台获取。获取方式请参考[获取项目ID](CloudArtifact_api_0015.xml)。 **约束限制**: 只能由英文字母、数字组成，且长度为32个字符。 **取值范围**: 不涉及。 **默认取值**: 无。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 是否返回默认仓库。 **约束限制**: 不涉及。 **取值范围**: true or false。 **默认取值**: false。
	Default *bool `json:"default,omitempty"`

	// **参数解释**: 仓库类型：snapshot 或 release。 **约束限制**: 不涉及。 **取值范围**: snapshot or releases。 **默认取值**: 不涉及。
	Policy *string `json:"policy,omitempty"`

	// **参数解释**: 仓库id，多个仓库id用英文逗号间隔。仓库id格式为{region}_{domainId}_{format}_{sequence}。可以从私有依赖库首页->仓库概览->仓库地址 url 中获取，最后两个\"/\"中间的字符串即为仓库id。 **约束限制**: 不涉及。 **取值范围**: 最大长度512。 **默认取值**: 不涉及。
	RepoIds *string `json:"repo_ids,omitempty"`

	// **参数解释**: 权限过滤设置，允许过滤读(r)和读写(rw)权限。 **约束限制**: 不涉及。 **取值范围**: r or rw。 **默认取值**: r。
	Access *string `json:"access,omitempty"`
}

func (o ListMavenListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMavenListRequest struct{}"
	}

	return strings.Join([]string{"ListMavenListRequest", string(data)}, " ")
}
