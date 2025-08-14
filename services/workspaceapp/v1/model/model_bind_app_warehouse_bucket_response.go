package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindAppWarehouseBucketResponse Response Object
type BindAppWarehouseBucketResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BindAppWarehouseBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindAppWarehouseBucketResponse struct{}"
	}

	return strings.Join([]string{"BindAppWarehouseBucketResponse", string(data)}, " ")
}
