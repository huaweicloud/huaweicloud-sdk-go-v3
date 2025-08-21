package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiffLinesResponse Response Object
type ShowDiffLinesResponse struct {

	// 文本信息
	Text           *string `json:"text,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDiffLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiffLinesResponse struct{}"
	}

	return strings.Join([]string{"ShowDiffLinesResponse", string(data)}, " ")
}
