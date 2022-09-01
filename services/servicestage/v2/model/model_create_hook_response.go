package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateHookResponse struct {

	// hook ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// hook类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 回滚URL。
	CallbackUrl    *string `json:"callback_url,omitempty" xml:"callback_url"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHookResponse struct{}"
	}

	return strings.Join([]string{"CreateHookResponse", string(data)}, " ")
}
