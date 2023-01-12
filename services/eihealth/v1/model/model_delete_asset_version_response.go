package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAssetVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetVersionResponse", string(data)}, " ")
}
