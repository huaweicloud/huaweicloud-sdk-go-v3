package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetResourceResponse Response Object
type CreateAssetResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAssetResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetResourceResponse", string(data)}, " ")
}
