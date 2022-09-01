package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBranchesByRepositoryIdResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *BranchList `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBranchesByRepositoryIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchesByRepositoryIdResponse struct{}"
	}

	return strings.Join([]string{"ShowBranchesByRepositoryIdResponse", string(data)}, " ")
}
