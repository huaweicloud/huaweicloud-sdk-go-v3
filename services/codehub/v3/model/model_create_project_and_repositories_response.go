package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectAndRepositoriesResponse Response Object
type CreateProjectAndRepositoriesResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *ProjectRepository `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProjectAndRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectAndRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectAndRepositoriesResponse", string(data)}, " ")
}
