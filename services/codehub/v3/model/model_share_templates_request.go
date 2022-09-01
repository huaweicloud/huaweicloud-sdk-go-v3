package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShareTemplatesRequest struct {

	// 仓库id
	RepositoryUuid string `json:"repository_uuid" xml:"repository_uuid"`

	Body *RepositoryTemplateVo `json:"body,omitempty" xml:"body"`
}

func (o ShareTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ShareTemplatesRequest", string(data)}, " ")
}
