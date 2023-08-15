package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessageEmailConfigResponse Response Object
type ShowMessageEmailConfigResponse struct {

	// 服务器地址
	Server *string `json:"server,omitempty"`

	// 展示名
	SubjectPrefix *string `json:"subject_prefix,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	Language       *LanguageEnum `json:"language,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowMessageEmailConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageEmailConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowMessageEmailConfigResponse", string(data)}, " ")
}
