package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareCacheGroupsRecord struct {

	// **参数解释：** 关联域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName *string `json:"domain_name,omitempty"`
}

func (o ShareCacheGroupsRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareCacheGroupsRecord struct{}"
	}

	return strings.Join([]string{"ShareCacheGroupsRecord", string(data)}, " ")
}
