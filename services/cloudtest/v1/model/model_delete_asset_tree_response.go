package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetTreeResponse Response Object
type DeleteAssetTreeResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAssetTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetTreeResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetTreeResponse", string(data)}, " ")
}
