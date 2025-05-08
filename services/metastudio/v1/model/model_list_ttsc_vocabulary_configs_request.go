package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTtscVocabularyConfigsRequest Request Object
type ListTtscVocabularyConfigsRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 自定义读法类型 CHINESE_G2P：拼音
	Type *string `json:"type,omitempty"`

	// 声音模型名称
	TtsServiceName *string `json:"tts_service_name,omitempty"`

	// 是否应用词表配置，从周边服务传递
	IsVocabularyConfigEnable *string `json:"is_vocabulary_config_enable,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 起始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。格式遵循：RFC 3339 如\"2021-01-10T10:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	// 搜索条件
	SearchKey *string `json:"search_key,omitempty"`
}

func (o ListTtscVocabularyConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTtscVocabularyConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListTtscVocabularyConfigsRequest", string(data)}, " ")
}
