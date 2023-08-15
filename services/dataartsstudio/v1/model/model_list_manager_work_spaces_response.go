package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagerWorkSpacesResponse Response Object
type ListManagerWorkSpacesResponse struct {

	// 当前工作空间用户记录数
	Count *int32 `json:"count,omitempty"`

	// 查询结果总页数
	TotalPage *int32 `json:"total_page,omitempty"`

	// 工作空间列表
	Data           *[]Workspacebody `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListManagerWorkSpacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagerWorkSpacesResponse struct{}"
	}

	return strings.Join([]string{"ListManagerWorkSpacesResponse", string(data)}, " ")
}
