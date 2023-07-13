package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTagToAssetResponse Response Object
type AddTagToAssetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddTagToAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagToAssetResponse struct{}"
	}

	return strings.Join([]string{"AddTagToAssetResponse", string(data)}, " ")
}
