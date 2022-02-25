package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteMaterialsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteMaterialsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMaterialsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteMaterialsResponse", string(data)}, " ")
}
