package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateMaterialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMaterialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMaterialResponse struct{}"
	}

	return strings.Join([]string{"UpdateMaterialResponse", string(data)}, " ")
}
