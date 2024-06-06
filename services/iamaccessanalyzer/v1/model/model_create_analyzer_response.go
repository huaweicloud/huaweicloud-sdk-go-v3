package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAnalyzerResponse Response Object
type CreateAnalyzerResponse struct {

	// 分析器的唯一标识符。
	Id *string `json:"id,omitempty"`

	// 分析器的唯一资源标识符。
	Urn            *string `json:"urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAnalyzerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnalyzerResponse struct{}"
	}

	return strings.Join([]string{"CreateAnalyzerResponse", string(data)}, " ")
}
