package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSmnTopicResponse Response Object
type AssociateSmnTopicResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o AssociateSmnTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSmnTopicResponse struct{}"
	}

	return strings.Join([]string{"AssociateSmnTopicResponse", string(data)}, " ")
}
