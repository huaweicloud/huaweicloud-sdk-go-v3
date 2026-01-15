package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomLineRequest Request Object
type ListCustomLineRequest struct {

	// 解析线路ID。
	LineId *string `json:"line_id,omitempty"`

	// 解析线路名称。
	Name *string `json:"name,omitempty"`

	// 分页查询时配置每页返回的资源个数。 当查询详细信息时：取值范围：0~100取值一般为10，20，50默认为100。 当查询概要信息时：取值范围：0~3000默认为3000。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当设置marker不为空时，以marker为分页起始标识，offset不生效。
	Offset *int32 `json:"offset,omitempty"`

	// 是否查询详细信息。  取值范围：  true：是，查询详细信息。 false：否，不查询详细信息。 默认为true。
	ShowDetail *bool `json:"show_detail,omitempty"`

	// **参数解释：** 资源状态。 **约束限制：** 不涉及。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ERROR：失败  **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// IP地址范围。
	Ip *string `json:"ip,omitempty"`
}

func (o ListCustomLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomLineRequest struct{}"
	}

	return strings.Join([]string{"ListCustomLineRequest", string(data)}, " ")
}
