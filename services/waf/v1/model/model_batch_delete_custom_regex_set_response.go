package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomRegexSetResponse Response Object
type BatchDeleteCustomRegexSetResponse struct {

	// 已成功删除的正则库ID列表
	Deleted *[]string `json:"deleted,omitempty"`

	// 删除失败的正则库列表
	Failed         *[]FailedCustomRegexSet `json:"failed,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchDeleteCustomRegexSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomRegexSetResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomRegexSetResponse", string(data)}, " ")
}
