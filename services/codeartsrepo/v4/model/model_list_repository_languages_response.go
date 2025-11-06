package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryLanguagesResponse Response Object
type ListRepositoryLanguagesResponse struct {

	// **参数解释：** 语言描述。
	Languages *[]LanguagesDto `json:"languages,omitempty"`

	// **参数解释：** 返回状态。
	Status *string `json:"status,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryLanguagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryLanguagesResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryLanguagesResponse", string(data)}, " ")
}
