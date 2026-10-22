package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemLexiconsResponse Response Object
type ListSystemLexiconsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 词典列表
	Items *[]SystemLexiconItem `json:"items,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSystemLexiconsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemLexiconsResponse struct{}"
	}

	return strings.Join([]string{"ListSystemLexiconsResponse", string(data)}, " ")
}
