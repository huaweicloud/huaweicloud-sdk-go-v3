package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBranchesByRepositoryIdResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *BranchList `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBranchesByRepositoryIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchesByRepositoryIdResponse struct{}"
	}

	return strings.Join([]string{"ShowBranchesByRepositoryIdResponse", string(data)}, " ")
}
