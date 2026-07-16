package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHyperinstanceTagsResponse Response Object
type CreateHyperinstanceTagsResponse struct {

	// 标签列表。
	Tags *[]TmsTag `json:"tags,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateHyperinstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHyperinstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateHyperinstanceTagsResponse", string(data)}, " ")
}
