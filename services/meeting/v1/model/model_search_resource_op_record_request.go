package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchResourceOpRecordRequest struct {

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

	// 查询过期时间在该时间戳之后的订单操作记录。
	StartExpireDate *int64 `json:"startExpireDate,omitempty"`

	// 查询过期时间在该时间戳之前的订单操作记录。
	EndExpireDate *int64 `json:"endExpireDate,omitempty"`

	// 查询操作时间在该时间戳之后的订单操作记录。
	StartOperateDate *int64 `json:"startOperateDate,omitempty"`

	// 查询操作时间在该时间戳之前的订单操作记录。
	EndOperateDate *int64 `json:"endOperateDate,omitempty"`

	// 订单资源类型。
	Type *string `json:"type,omitempty"`

	// 当前仅当资源类型为vmr时生效。
	TypeId *string `json:"typeId,omitempty"`

	// 操作类型。 - 0：添加 - 1：删除 - 2：修改 - 3：停用 - 4：启用
	OperateType *int32 `json:"operateType,omitempty"`
}

func (o SearchResourceOpRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceOpRecordRequest struct{}"
	}

	return strings.Join([]string{"SearchResourceOpRecordRequest", string(data)}, " ")
}
