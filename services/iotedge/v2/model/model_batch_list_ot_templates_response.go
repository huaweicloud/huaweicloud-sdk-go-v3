package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListOtTemplatesResponse Response Object
type BatchListOtTemplatesResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 每页记录数
	Templates      *[]QueryOtTemplateBriefRespDto `json:"templates,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o BatchListOtTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListOtTemplatesResponse struct{}"
	}

	return strings.Join([]string{"BatchListOtTemplatesResponse", string(data)}, " ")
}
