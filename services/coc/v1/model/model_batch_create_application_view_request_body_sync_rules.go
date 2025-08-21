package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBodySyncRules struct {

	// **参数解释：** 企业项目id。 **约束限制：** 不涉及。 **取值范围：** 请选择[[企业管理](https://support.huaweicloud.com/usermanual-em/em_eps_qs_0400.html)](tag:hws)[[企业管理](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/em_eps_qs_0400.html)](tag:hws_hk)中存在的项目ID。 **默认取值：** 不涉及。
	EpId *string `json:"ep_id,omitempty"`

	// **参数解释：** 关联标签。 **约束限制：** 不涉及。 **取值范围：** 可自定义。 **默认取值：** 不涉及。
	RuleTags *string `json:"rule_tags,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodySyncRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodySyncRules struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodySyncRules", string(data)}, " ")
}
