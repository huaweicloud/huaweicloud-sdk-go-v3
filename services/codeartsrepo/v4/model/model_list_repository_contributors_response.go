package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryContributorsResponse Response Object
type ListRepositoryContributorsResponse struct {

	// 贡献者信息列表
	Body *[]ContributorDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryContributorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryContributorsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryContributorsResponse", string(data)}, " ")
}
