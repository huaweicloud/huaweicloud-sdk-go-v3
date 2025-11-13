package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWebsiteRecordSetRequest Request Object
type ShowWebsiteRecordSetRequest struct {

	// **参数解释：** 域名ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId string `json:"zone_id"`

	// **参数解释：** 分页查询时配置每页返回的资源个数。 **约束限制：** 不涉及。 **取值范围：** 0~500。 **默认取值：** 500
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。 **约束限制：** 当设置marker不为空时，以marker为分页起始标识，offset不生效。 **取值范围：** 0~2147483647。 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowWebsiteRecordSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebsiteRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowWebsiteRecordSetRequest", string(data)}, " ")
}
