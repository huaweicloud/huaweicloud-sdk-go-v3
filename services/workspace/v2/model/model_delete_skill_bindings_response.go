package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSkillBindingsResponse Response Object
type DeleteSkillBindingsResponse struct {

	// 请求总数。
	Total *int32 `json:"total,omitempty"`

	// 成功数量。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails *[]BindingFailedDetail `json:"failed_details,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSkillBindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSkillBindingsResponse struct{}"
	}

	return strings.Join([]string{"DeleteSkillBindingsResponse", string(data)}, " ")
}
