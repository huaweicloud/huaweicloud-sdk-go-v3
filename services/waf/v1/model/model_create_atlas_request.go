package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAtlasRequest Request Object
type CreateAtlasRequest struct {
	Body *AiguardAtlasBody `json:"body,omitempty"`
}

func (o CreateAtlasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAtlasRequest struct{}"
	}

	return strings.Join([]string{"CreateAtlasRequest", string(data)}, " ")
}
