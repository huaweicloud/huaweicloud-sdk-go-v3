package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLinksResponse Response Object
type ListLinksResponse struct {

	// 可用链路信息。
	JobLinks *[]JobLinkResp `json:"job_links,omitempty"`

	// 可用链路总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLinksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLinksResponse struct{}"
	}

	return strings.Join([]string{"ListLinksResponse", string(data)}, " ")
}
