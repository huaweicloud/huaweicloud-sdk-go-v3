package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryForksResponse Response Object
type ListRepositoryForksResponse struct {

	// 仓库信息
	Body           *[]ForkRepositoryDto `json:"body,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListRepositoryForksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryForksResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryForksResponse", string(data)}, " ")
}
