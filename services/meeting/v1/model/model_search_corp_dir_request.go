package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SearchCorpDirRequest struct {
	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成

	XRequestId *string `json:"X-Request-Id,omitempty"`
	// 语言参数，默认为中文zh-CN, 英文为en-US

	AcceptLanguage *string `json:"Accept-Language,omitempty"`
	// 查询偏移量,若超过最大数量，则返回最后一页的数据 默认值：0

	Offset *int32 `json:"offset,omitempty"`
	// 查询数量 默认值：0

	Limit *int32 `json:"limit,omitempty"`
	// 搜索条件。支持账号、姓名、手机、邮箱模糊搜索

	SearchKey *string `json:"searchKey,omitempty"`
	// 部门编码 maxLength：32 minLength：0

	DeptCode *string `json:"deptCode,omitempty"`
	// 是否查询子部门下的用户 默认值：true

	QuerySubDept *bool `json:"querySubDept,omitempty"`
	// * 搜索范围 * NORMAL_USER表示查询普通用户。返回普通用户（响应中isHardTerminal=false，type=NORMAL_USER） * HARD_TERMINAL表示查询硬终端用户。返回大屏用户（响应中isHardTerminal=false，type=WHITE_BOARD）和硬终端用户（响应中isHardTerminal=true，type=HARD_TERMINAL） * ALL表示查询所有用户。 * 默认值为ALL

	SearchScope *string `json:"searchScope,omitempty"`
}

func (o SearchCorpDirRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpDirRequest struct{}"
	}

	return strings.Join([]string{"SearchCorpDirRequest", string(data)}, " ")
}
