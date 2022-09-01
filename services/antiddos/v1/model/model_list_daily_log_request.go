package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDailyLogRequest struct {

	// 用户EIP对应的ID
	FloatingIpId string `json:"floating_ip_id" xml:"floating_ip_id"`

	// 可选范围： - desc：表示时间降序 - asc：表示时间升序 默认值为“desc”。
	SortDir *string `json:"sort_dir,omitempty" xml:"sort_dir"`

	// 返回结果个数限制，此次查询返回数量最大值，取值范围：1～100，与offset配合使用。 若“limit”与“offset”均不携带则返回所有主机列表。
	Limit *string `json:"limit,omitempty" xml:"limit"`

	// 偏移量，“limit”携带时此字段有效。
	Offset *string `json:"offset,omitempty" xml:"offset"`

	// 用户EIP
	Ip *string `json:"ip,omitempty" xml:"ip"`
}

func (o ListDailyLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDailyLogRequest struct{}"
	}

	return strings.Join([]string{"ListDailyLogRequest", string(data)}, " ")
}
