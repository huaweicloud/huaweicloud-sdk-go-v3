package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddMaterialResponse Response Object
type AddMaterialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddMaterialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMaterialResponse struct{}"
	}

	return strings.Join([]string{"AddMaterialResponse", string(data)}, " ")
}
