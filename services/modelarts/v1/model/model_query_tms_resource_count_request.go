package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryTmsResourceCountRequest 查询TMS资源总量请求体
type QueryTmsResourceCountRequest struct {

	// **参数解释：** 匹配项，目前只支持资源名称的模糊匹配。 **约束限制：** 不涉及。
	Matches *[]TmsMatch `json:"matches,omitempty"`

	// **参数解释：** 标签匹配项，只支持多个标签与操作，不携带表示查询所有资源。 **约束限制：** 不涉及。
	Tags *[]CombineInferTmsTags `json:"tags,omitempty"`

	// **参数解释：** 是否只查询没有打标签的资源。 **约束限制：** 不涉及。 **取值范围：** true：只查询没有打标签的资源。 false：查询所有资源。 **默认取值：** 不涉及。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
}

func (o QueryTmsResourceCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTmsResourceCountRequest struct{}"
	}

	return strings.Join([]string{"QueryTmsResourceCountRequest", string(data)}, " ")
}
