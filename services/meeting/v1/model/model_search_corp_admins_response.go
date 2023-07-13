package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpAdminsResponse Response Object
type SearchCorpAdminsResponse struct {

	// 页面起始页，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 总数量。
	Count *int32 `json:"count,omitempty"`

	// 查询企业管理员返回的结果。
	Data           *[]QueryCorpAdminResultDto `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o SearchCorpAdminsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpAdminsResponse struct{}"
	}

	return strings.Join([]string{"SearchCorpAdminsResponse", string(data)}, " ")
}
