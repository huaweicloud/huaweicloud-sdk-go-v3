package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCertTagsResponse Response Object
type BatchDeleteCertTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteCertTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCertTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteCertTagsResponse", string(data)}, " ")
}
