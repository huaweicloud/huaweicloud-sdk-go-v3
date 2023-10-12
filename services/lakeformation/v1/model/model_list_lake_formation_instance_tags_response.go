package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLakeFormationInstanceTagsResponse Response Object
type ListLakeFormationInstanceTagsResponse struct {

	// 标签
	Tags *[]ResourceTagParam `json:"tags,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLakeFormationInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLakeFormationInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListLakeFormationInstanceTagsResponse", string(data)}, " ")
}
