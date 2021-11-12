package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAuditlogLinksResponse struct {
	// 审计日志下载链接列表，有效时间5分钟。

	Links          *[]string `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAuditlogLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogLinksResponse struct{}"
	}

	return strings.Join([]string{"ListAuditlogLinksResponse", string(data)}, " ")
}
