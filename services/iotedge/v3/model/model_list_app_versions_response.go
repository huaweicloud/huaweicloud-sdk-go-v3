package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppVersionsResponse Response Object
type ListAppVersionsResponse struct {

	// 总记录数
	Count *int32 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 每页记录数
	AppVersions    *[]QueryAppVersionResponseDto `json:"app_versions,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListAppVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListAppVersionsResponse", string(data)}, " ")
}
