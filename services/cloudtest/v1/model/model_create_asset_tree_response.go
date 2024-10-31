package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetTreeResponse Response Object
type CreateAssetTreeResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetTreeResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetTreeResponse", string(data)}, " ")
}
