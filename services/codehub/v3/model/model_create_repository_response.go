package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRepositoryResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *Repository `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepositoryResponse struct{}"
	}

	return strings.Join([]string{"CreateRepositoryResponse", string(data)}, " ")
}
