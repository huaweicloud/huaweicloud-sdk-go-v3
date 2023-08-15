package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimTemplateMaterialsResponseMode 查询素材返回体。
type ListAimTemplateMaterialsResponseMode struct {
	PageInfo *PageInfo `json:"page_info"`

	// 模板素材列表。
	Results *[]Material `json:"results,omitempty"`
}

func (o ListAimTemplateMaterialsResponseMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplateMaterialsResponseMode struct{}"
	}

	return strings.Join([]string{"ListAimTemplateMaterialsResponseMode", string(data)}, " ")
}
