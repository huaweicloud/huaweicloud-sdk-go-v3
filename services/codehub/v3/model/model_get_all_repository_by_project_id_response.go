package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GetAllRepositoryByProjectIdResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepoListInfo `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetAllRepositoryByProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAllRepositoryByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"GetAllRepositoryByProjectIdResponse", string(data)}, " ")
}
