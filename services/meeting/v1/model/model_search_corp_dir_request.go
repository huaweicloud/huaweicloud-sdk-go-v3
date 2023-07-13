package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpDirRequest Request Object
type SearchCorpDirRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页的数据。 默认值：0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件。支持帐号、SIP号码、名称、手机、邮箱模糊搜索。
	SearchKey *string `json:"searchKey,omitempty"`

	// 部门编码。
	DeptCode *string `json:"deptCode,omitempty"`

	// 是否查询子部门下的用户。 默认值：true。
	QuerySubDept *bool `json:"querySubDept,omitempty"`

	// 搜索范围。默认值为ALL。 * NORMAL_USER - 查询普通用户。返回普通用户（响应中isHardTerminal=false，type=NORMAL_USER） * HARD_TERMINAL - 查询硬终端用户。返回大屏用户（响应中isHardTerminal=false，type=WHITE_BOARD）和硬终端用户（响应中isHardTerminal=true，type=HARD_TERMINAL） * ALL - 查询所有用户。
	SearchScope *string `json:"searchScope,omitempty"`
}

func (o SearchCorpDirRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpDirRequest struct{}"
	}

	return strings.Join([]string{"SearchCorpDirRequest", string(data)}, " ")
}
