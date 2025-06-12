package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryResponse Response Object
type ListRepositoryResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *RepositoriesResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryResponse", string(data)}, " ")
}
