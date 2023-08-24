package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListAppConfigsTemplatesResponse Response Object
type BatchListAppConfigsTemplatesResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 模板列表
	Templates      *[]QueryAppConfigsTemplateBriefRespDto `json:"templates,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o BatchListAppConfigsTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListAppConfigsTemplatesResponse struct{}"
	}

	return strings.Join([]string{"BatchListAppConfigsTemplatesResponse", string(data)}, " ")
}
