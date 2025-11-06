package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryReviewsResponse Response Object
type ListRepositoryReviewsResponse struct {
	Body           *[]ReviewDto `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRepositoryReviewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryReviewsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryReviewsResponse", string(data)}, " ")
}
