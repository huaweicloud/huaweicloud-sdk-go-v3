package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetTreeResponse Response Object
type UpdateAssetTreeResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAssetTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetTreeResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssetTreeResponse", string(data)}, " ")
}
