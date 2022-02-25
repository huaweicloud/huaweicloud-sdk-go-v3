package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchCorpResourcesRequest struct {
	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成

	XRequestId *string `json:"X-Request-Id,omitempty"`
	// 语言参数，默认为中文zh-CN, 英文为en-US

	AcceptLanguage *string `json:"Accept-Language,omitempty"`
	// 查询偏移量,若超过最大数量，则返回最后一页的数据 默认值：0

	Offset *int32 `json:"offset,omitempty"`
	// 查询数量 默认值：0

	Limit *int32 `json:"limit,omitempty"`
	// 搜索条件，支持resourceId模糊查询。

	SearchKey *string `json:"searchKey,omitempty"`
	// 查询过期时间在该时间戳之后的资源项

	StartExpireDate *int64 `json:"startExpireDate,omitempty"`
	// 查询过期时间在该时间戳之前的资源项

	EndExpireDate *int64 `json:"endExpireDate,omitempty"`
	// 资源类型。 - VMR        - 云会议室 - CONF_CALL  - 会议并发数 - HARD_1080P - 1080P硬终端 - HARD_720P  - 720P硬终端 - SOFT       - 软终端用户数 - ROOM       - 大屏软终端 - LIVE       - 直播推流 - RECORD     - 录播空间 - HARD_THIRD_PARTY - 第三方硬终端账号 - HUAWEI_VISION -智慧屏 说明：查询网络研讨会资源时type字段需上送VMR

	Type *string `json:"type,omitempty"`
	// VMR模式，type为vmr时传递该参数 * 0：个人会议ID * 1：云会议室 * 2：网络研讨会

	VmrMode *int32 `json:"vmrMode,omitempty"`
	// 资源类型Id,若想搜索5方VMR时，需携带5方vmrpkg对应的id

	TypeId *string `json:"typeId,omitempty"`
	// 订单Id

	OrderId *string `json:"orderId,omitempty"`
	// 订单状态。 - 0：正常 - 1：到期 - 2：停用

	Status *int32 `json:"status,omitempty"`
}

func (o SearchCorpResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpResourcesRequest struct{}"
	}

	return strings.Join([]string{"SearchCorpResourcesRequest", string(data)}, " ")
}
