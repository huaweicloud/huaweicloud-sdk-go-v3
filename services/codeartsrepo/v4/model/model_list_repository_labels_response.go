package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryLabelsResponse Response Object
type ListRepositoryLabelsResponse struct {

	// 标签列表
	Body           *[]LabelBasicDto `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListRepositoryLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryLabelsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryLabelsResponse", string(data)}, " ")
}
