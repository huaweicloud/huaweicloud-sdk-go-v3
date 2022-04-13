package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBranchesByTwoRepositoryIdResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *TagList `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBranchesByTwoRepositoryIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchesByTwoRepositoryIdResponse struct{}"
	}

	return strings.Join([]string{"ShowBranchesByTwoRepositoryIdResponse", string(data)}, " ")
}
