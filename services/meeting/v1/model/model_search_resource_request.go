package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResourceRequest Request Object
type SearchResourceRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页的数据。 默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件。
	SearchKey *string `json:"searchKey,omitempty"`

	// 企业id。
	CorpId string `json:"corp_id"`

	// 查询过期时间在该时间戳之后的资源项。
	StartExpireDate *int64 `json:"startExpireDate,omitempty"`

	// 查询过期时间在该时间戳之前的资源项。
	EndExpireDate *int64 `json:"endExpireDate,omitempty"`

	// 资源类型。
	Type *string `json:"type,omitempty"`

	// 资源类型Id,若想搜索5方VMR时，需携带5方vmrpkg对应的id。
	TypeId *string `json:"typeId,omitempty"`

	// 订单状态: - 0：正常 - 1：到期，仅查询时返回 - 2：停用
	Status *int32 `json:"status,omitempty"`
}

func (o SearchResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceRequest struct{}"
	}

	return strings.Join([]string{"SearchResourceRequest", string(data)}, " ")
}
