package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSignatureDetailsRequest struct {

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 签名ID
	SignatureId *string `json:"signature_id,omitempty"`

	// 签名名称
	SignatureName *string `json:"signature_name,omitempty"`

	// 签名类型。支持枚举值： 1. VERIFY_CODE_TYPE: 验证码类 2. PROMOTION_TYPE: 推广类 3. NOTIFY_TYPE: 通知类
	SignatureType *string `json:"signature_type,omitempty"`

	// 地域
	Site *string `json:"site,omitempty"`

	// 排序方式 1. desc：降序 2. asc：升序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 状态 1. PENDING_REVIEW：待审核 2. PROCESSING：内容审核通过，签名处理中 3. REVIEW_PASSED：处理完毕 4. REVIEW_NOT_PASS：审核不通过 5. TO_BE_ACTIVATED：待激活 6. PENDING_ACTIVATE：激活审核中
	Status *string `json:"status,omitempty"`
}

func (o ListSignatureDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSignatureDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSignatureDetailsRequest", string(data)}, " ")
}
