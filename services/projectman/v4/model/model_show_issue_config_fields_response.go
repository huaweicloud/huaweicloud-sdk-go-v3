package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIssueConfigFieldsResponse Response Object
type ShowIssueConfigFieldsResponse struct {

	// 状态码
	Status *string `json:"status,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	Result         *IssueConfigFieldsResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowIssueConfigFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueConfigFieldsResponse struct{}"
	}

	return strings.Join([]string{"ShowIssueConfigFieldsResponse", string(data)}, " ")
}
