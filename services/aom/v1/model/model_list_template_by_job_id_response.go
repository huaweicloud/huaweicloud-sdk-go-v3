package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTemplateByJobIdResponse struct {

	// 总数
	TotalElements *int64 `json:"total_elements,omitempty"`

	// 查询作业信息集合
	Elements       *[]Template `json:"elements,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListTemplateByJobIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateByJobIdResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateByJobIdResponse", string(data)}, " ")
}
