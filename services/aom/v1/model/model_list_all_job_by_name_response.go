package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllJobByNameResponse Response Object
type ListAllJobByNameResponse struct {

	// 总数。
	TotalElements *int64 `json:"total_elements,omitempty"`

	// 查询作业信息集合。
	Elements       *[]Job `json:"elements,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAllJobByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllJobByNameResponse struct{}"
	}

	return strings.Join([]string{"ListAllJobByNameResponse", string(data)}, " ")
}
