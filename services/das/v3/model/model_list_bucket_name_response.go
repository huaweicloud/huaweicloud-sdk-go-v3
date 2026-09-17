package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBucketNameResponse Response Object
type ListBucketNameResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBucketNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBucketNameResponse struct{}"
	}

	return strings.Join([]string{"ListBucketNameResponse", string(data)}, " ")
}
