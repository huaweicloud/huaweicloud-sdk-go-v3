package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPubInfosRequest struct {

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`

	// 服务号状态。  - 1：未生效  - 2：已生效  - 3：已失效  - 4：已冻结
	State *int32 `json:"state,omitempty"`

	// 开始上线时间，格式为：yyyy-MM-ddTHH:mm:ssZ。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 结束上线时间，格式为：yyyy-MM-ddTHH:mm:ssZ。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 服务号所属行业。 - 1：金融理财 - 2：社交通讯 - 3：影音娱乐 - 4：旅游出行 - 5：购物 - 6：本地生活 - 7：运动健康 - 8：教育培训 - 9：新闻阅读 - 10：运营商  - 11：其他
	Industry *string `json:"industry,omitempty"`

	// 审核状态。  - 1：审核中 - 2：审核通过 - 3：驳回
	ApproveState *int32 `json:"approve_state,omitempty"`
}

func (o ListPubInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPubInfosRequest struct{}"
	}

	return strings.Join([]string{"ListPubInfosRequest", string(data)}, " ")
}
