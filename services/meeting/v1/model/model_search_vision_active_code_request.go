package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchVisionActiveCodeRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id"`

	// 语言参数，默认为中文zh-CN, 英文为en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty" xml:"Accept-Language"`

	// 查询偏移量,若超过最大数量，则返回最后一页的数据 默认值：0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询数量 默认值：0
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 搜索条件,支持激活码、终端名称模糊搜索。
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey"`

	// 终端类型 - idea-hub:智能协作大屏 - huawei-vision:智慧屏TV - welink-desktop(iwb):电子白板
	DevType *string `json:"devType,omitempty" xml:"devType"`
}

func (o SearchVisionActiveCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchVisionActiveCodeRequest struct{}"
	}

	return strings.Join([]string{"SearchVisionActiveCodeRequest", string(data)}, " ")
}
