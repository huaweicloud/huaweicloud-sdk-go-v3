package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomAnswersResponse Response Object
type BatchDeleteCustomAnswersResponse struct {

	// 已删除的代答库ID列表
	Deleted *[]string `json:"deleted,omitempty"`

	// 删除失败的代答库列表
	Failed         *[]FailedCustomAnswers `json:"failed,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o BatchDeleteCustomAnswersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomAnswersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomAnswersResponse", string(data)}, " ")
}
