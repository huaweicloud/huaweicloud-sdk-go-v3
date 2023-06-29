package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetRepositoryByProjectIdResponse Response Object
type GetRepositoryByProjectIdResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepoInfo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetRepositoryByProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetRepositoryByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"GetRepositoryByProjectIdResponse", string(data)}, " ")
}
