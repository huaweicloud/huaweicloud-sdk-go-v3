package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWeakPasswordRequest Request Object
type CheckWeakPasswordRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *WeakPasswordCheckRequestBody `json:"body,omitempty"`
}

func (o CheckWeakPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeakPasswordRequest struct{}"
	}

	return strings.Join([]string{"CheckWeakPasswordRequest", string(data)}, " ")
}
