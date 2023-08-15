package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateCertTagsResponse Response Object
type BatchCreateCertTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateCertTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCertTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateCertTagsResponse", string(data)}, " ")
}
