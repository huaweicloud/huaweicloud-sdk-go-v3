package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBranchesByTwoRepositoryIdResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *TagList `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBranchesByTwoRepositoryIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchesByTwoRepositoryIdResponse struct{}"
	}

	return strings.Join([]string{"ShowBranchesByTwoRepositoryIdResponse", string(data)}, " ")
}
