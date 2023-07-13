package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRootsResponse Response Object
type ListRootsResponse struct {

	// 在组织中定义的根列表。
	Roots *[]RootDto `json:"roots,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRootsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRootsResponse struct{}"
	}

	return strings.Join([]string{"ListRootsResponse", string(data)}, " ")
}
