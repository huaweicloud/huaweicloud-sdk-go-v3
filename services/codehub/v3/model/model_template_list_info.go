package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateListInfo struct {

	// 仓库列表
	Repos *[]DevstarRepoInfo `json:"repos,omitempty"`

	// 仓库总数
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o TemplateListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateListInfo struct{}"
	}

	return strings.Join([]string{"TemplateListInfo", string(data)}, " ")
}
