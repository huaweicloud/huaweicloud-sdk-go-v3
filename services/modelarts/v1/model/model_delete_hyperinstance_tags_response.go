package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHyperinstanceTagsResponse Response Object
type DeleteHyperinstanceTagsResponse struct {

	// 标签列表。
	Tags *[]TmsTag `json:"tags,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHyperinstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHyperinstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteHyperinstanceTagsResponse", string(data)}, " ")
}
