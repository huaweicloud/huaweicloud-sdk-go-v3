package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
