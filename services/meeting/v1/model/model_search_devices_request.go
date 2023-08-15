package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchDevicesRequest Request Object
type SearchDevicesRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量，若超过最大数量，则返回最后一页。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件。支持名称、SN模糊查询。
	SearchKey *string `json:"searchKey,omitempty"`

	// 终端型号，枚举类型。当前支持TE系列硬件终端，具体的终端类型可以通过获取所有终端类型接口查询。
	Model *string `json:"model,omitempty"`

	// 部门编码，默认为根部门。 默认值：1。
	DeptCode *string `json:"deptCode,omitempty"`

	// 是否查询子部门。 默认值：true。
	EnableSubDept *bool `json:"enableSubDept,omitempty"`
}

func (o SearchDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDevicesRequest struct{}"
	}

	return strings.Join([]string{"SearchDevicesRequest", string(data)}, " ")
}
