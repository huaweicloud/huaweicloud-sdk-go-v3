package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShareCacheGroupsRequstBody **参数解释：** 共享缓存组配置 **约束限制：** 不涉及
type CreateShareCacheGroupsRequstBody struct {

	// **参数解释：** 共享缓存组名称 **约束限制：** 不涉及 **取值范围：** - 1-128个字符 - 不支持特殊字符“%”、“&”、“=”、“?”、“$”\"、“\"”、“<”、“>” **默认取值：** 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释：** 主域名 **约束限制：** 只能有一个主域名 **取值范围：** 不涉及 **默认取值：** 不涉及
	PrimaryDomain *string `json:"primary_domain,omitempty"`

	// **参数解释：** 关联域名 **约束限制：** - 必须传入主域名 - 最多可包含50个关联域名
	ShareCacheRecords *[]ShareCacheGroupsRecord `json:"share_cache_records,omitempty"`
}

func (o CreateShareCacheGroupsRequstBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareCacheGroupsRequstBody struct{}"
	}

	return strings.Join([]string{"CreateShareCacheGroupsRequstBody", string(data)}, " ")
}
