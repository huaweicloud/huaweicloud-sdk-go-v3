package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppWarehouseBucketRequest Request Object
type ShowAppWarehouseBucketRequest struct {
}

func (o ShowAppWarehouseBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppWarehouseBucketRequest struct{}"
	}

	return strings.Join([]string{"ShowAppWarehouseBucketRequest", string(data)}, " ")
}
