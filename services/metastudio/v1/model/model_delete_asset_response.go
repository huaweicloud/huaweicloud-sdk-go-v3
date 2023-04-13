package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAssetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetResponse", string(data)}, " ")
}
