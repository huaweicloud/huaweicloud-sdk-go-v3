package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepositorySystemLabelsResponse Response Object
type CreateRepositorySystemLabelsResponse struct {

	// 标签列表
	Body           *[]LabelDetailDto `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateRepositorySystemLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepositorySystemLabelsResponse struct{}"
	}

	return strings.Join([]string{"CreateRepositorySystemLabelsResponse", string(data)}, " ")
}
