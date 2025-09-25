package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityDataCategoriesResponse Response Object
type BatchDeleteSecurityDataCategoriesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecurityDataCategoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityDataCategoriesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityDataCategoriesResponse", string(data)}, " ")
}
