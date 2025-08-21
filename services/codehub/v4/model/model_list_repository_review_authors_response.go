package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryReviewAuthorsResponse Response Object
type ListRepositoryReviewAuthorsResponse struct {
	Body           *[]UserBasicDto `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListRepositoryReviewAuthorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryReviewAuthorsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryReviewAuthorsResponse", string(data)}, " ")
}
