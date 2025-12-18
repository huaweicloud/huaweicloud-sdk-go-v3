package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicAccess struct {

	// **参数解释：** 允许访问集群API的白名单网段列表，建议对VPC网段、容器网段放通。 **约束限制：** 该字段仅支持创建集群时传入，更新时指定无效 **取值范围：** 不涉及 **默认取值：** 默认无白名单配置，为[\"0.0.0.0/0\"]。
	Cidrs *[]string `json:"cidrs,omitempty"`
}

func (o PublicAccess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicAccess struct{}"
	}

	return strings.Join([]string{"PublicAccess", string(data)}, " ")
}
