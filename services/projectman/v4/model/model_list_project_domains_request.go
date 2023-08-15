package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectDomainsRequest Request Object
type ListProjectDomainsRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 一次返回的数据,最小1,最大100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListProjectDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectDomainsRequest", string(data)}, " ")
}
