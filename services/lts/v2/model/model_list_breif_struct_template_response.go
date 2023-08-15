package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBreifStructTemplateResponse Response Object
type ListBreifStructTemplateResponse struct {

	// 结构化模板缩略信息列表
	Results        *[]BriefStructTemplateModel `json:"results,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListBreifStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBreifStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListBreifStructTemplateResponse", string(data)}, " ")
}
