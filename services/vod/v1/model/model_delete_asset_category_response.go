package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetCategoryResponse Response Object
type DeleteAssetCategoryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssetCategoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetCategoryResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetCategoryResponse", string(data)}, " ")
}
