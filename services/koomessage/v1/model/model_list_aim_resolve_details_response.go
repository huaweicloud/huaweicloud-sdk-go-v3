package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimResolveDetailsResponse Response Object
type ListAimResolveDetailsResponse struct {

	// 查询解析结果集。
	ResolveDetails *[]AimResolveDetail `json:"resolve_details,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimResolveDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimResolveDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListAimResolveDetailsResponse", string(data)}, " ")
}
