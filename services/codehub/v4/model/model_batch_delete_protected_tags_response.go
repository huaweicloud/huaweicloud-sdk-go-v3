package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProtectedTagsResponse Response Object
type BatchDeleteProtectedTagsResponse struct {

	// **参数解释：** 状态码。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteProtectedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedTagsResponse", string(data)}, " ")
}
