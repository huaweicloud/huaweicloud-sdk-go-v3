package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryHyperinstanceTagsResponse Response Object
type QueryHyperinstanceTagsResponse struct {

	// 标签列表。
	Tags *[]TmsTag `json:"tags,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o QueryHyperinstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryHyperinstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"QueryHyperinstanceTagsResponse", string(data)}, " ")
}
