package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindAppWarehouseBucketRequest Request Object
type BindAppWarehouseBucketRequest struct {
	Body *BucketNameReq `json:"body,omitempty"`
}

func (o BindAppWarehouseBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindAppWarehouseBucketRequest struct{}"
	}

	return strings.Join([]string{"BindAppWarehouseBucketRequest", string(data)}, " ")
}
