package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainTopPathRequest Request Object
type ListCdnDomainTopPathRequest struct {

	// **参数解释：** 查询起始时间戳 **约束限制：** 该参数只能传0点毫秒时间戳 **取值范围：** 不涉及 **默认取值：** 不涉及
	StartTime int64 `json:"start_time"`

	// **参数解释：** 查询结束时间戳 **约束限制：** 该参数只能传0点毫秒时间戳 **取值范围：** 不涉及 **默认取值：** 不涉及
	EndTime int64 `json:"end_time"`

	// **参数解释：** 域名列表 > 如果域名在查询时间段内无数据，结果将不返回该域名的信息  **约束限制：** 仅支持查询已经在CDN创建成功的域名 **取值范围：** - all表示查询名下全部域名 - 多个域名以逗号（半角）分隔，如：www.test1.com,www.test2.com **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 统计指标类型 **约束限制：** 不涉及 **取值范围：** - flux：流量 - req_num：请求数 **默认取值：** 不涉及
	StatType string `json:"stat_type"`

	// **参数解释：** 服务范围 **约束限制：** 不涉及 **取值范围：** - mainland_china：中国大陆 - outside_mainland_china：中国大陆境外 - global：全球 **默认取值：** global：全球
	ServiceArea *string `json:"service_area,omitempty"`

	// **参数解释：** 域名所属账号的domain_id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UserDomainId *string `json:"user_domain_id,omitempty"`
}

func (o ListCdnDomainTopPathRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainTopPathRequest struct{}"
	}

	return strings.Join([]string{"ListCdnDomainTopPathRequest", string(data)}, " ")
}
