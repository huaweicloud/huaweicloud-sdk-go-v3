package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAssetResponse Response Object
type ImportAssetResponse struct {
	Code *string `json:"code,omitempty"`

	Data *interface{} `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAssetResponse struct{}"
	}

	return strings.Join([]string{"ImportAssetResponse", string(data)}, " ")
}
