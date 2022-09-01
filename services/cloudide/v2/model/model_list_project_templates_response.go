package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectTemplatesResponse struct {

	// 模板列表
	Templates *[]ProjectTemplates `json:"templates,omitempty" xml:"templates"`

	// 状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTemplatesResponse", string(data)}, " ")
}
