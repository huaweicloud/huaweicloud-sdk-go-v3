package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllRepositoryByTwoProjectIdResponse Response Object
type ShowAllRepositoryByTwoProjectIdResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepoListInfoV2 `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAllRepositoryByTwoProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllRepositoryByTwoProjectIdResponse struct{}"
	}

	return strings.Join([]string{"ShowAllRepositoryByTwoProjectIdResponse", string(data)}, " ")
}
