package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplatesTwoRequest struct {
	// 仓库id

	RepositoryUuid string `json:"repository_uuid"`

	Body *RepositoryTemplateVo2 `json:"body,omitempty"`
}

func (o ListTemplatesTwoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesTwoRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesTwoRequest", string(data)}, " ")
}
