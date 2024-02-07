package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessSitesResponse Response Object
type ListAccessSitesResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 接入点列表
	AccessSites *[]ListAccessSites `json:"access_sites,omitempty"`

	PageInfo *ListGlobalEipsResponseBodyPageInfo `json:"page_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAccessSitesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessSitesResponse struct{}"
	}

	return strings.Join([]string{"ListAccessSitesResponse", string(data)}, " ")
}
