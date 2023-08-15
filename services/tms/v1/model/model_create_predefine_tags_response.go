package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePredefineTagsResponse Response Object
type CreatePredefineTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePredefineTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePredefineTagsResponse struct{}"
	}

	return strings.Join([]string{"CreatePredefineTagsResponse", string(data)}, " ")
}
