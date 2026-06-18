package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryNavigationLanguageResponse Response Object
type ShowRepositoryNavigationLanguageResponse struct {

	// **参数解释：** 结果标识。 **约束限制：** 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：** 结果消息。 **约束限制：** 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释：** 语言列表。 **约束限制：** 不涉及。
	LanguageList   *[]LanguageDto `json:"language_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowRepositoryNavigationLanguageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryNavigationLanguageResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryNavigationLanguageResponse", string(data)}, " ")
}
