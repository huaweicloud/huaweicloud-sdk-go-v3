package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryTemplatesResponse Response Object
type ListRepositoryTemplatesResponse struct {

	// 模板仓列表
	Body *[]RepositoryTemplateDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryTemplatesResponse", string(data)}, " ")
}
