package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetNewResponse Response Object
type DeleteAssetNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetNewResponse", string(data)}, " ")
}
