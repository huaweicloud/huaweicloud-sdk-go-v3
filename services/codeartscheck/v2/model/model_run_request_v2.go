package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunRequestV2 代码仓信息
type RunRequestV2 struct {

	// **参数解释：** 废弃参数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 废弃参数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	AccessToken *string `json:"access_token,omitempty"`

	// **参数解释：** 废弃参数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释：** 废弃参数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GitBranch *string `json:"git_branch,omitempty"`

	// **参数解释：** 代码仓拉取MR代码的方式，推荐在检查阶段拉取MR代码失败时使用此参数进行拉取方式的指定。 **约束限制：** 不涉及。 **取值范围：** - pull：以pull/xx/xx方式进行拉取。 - merge_request：以merge_request/xx/xx方式进行拉取。 **默认取值：** 后台自动识别代码仓类型以进行MR代码的拉取。 gitee仓：以pull/xx/xx方式进行拉取。 非gitee仓：以merge_request/xx/xx方式进行拉取。
	Ref *string `json:"ref,omitempty"`
}

func (o RunRequestV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunRequestV2 struct{}"
	}

	return strings.Join([]string{"RunRequestV2", string(data)}, " ")
}
