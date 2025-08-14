package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppWarehouseBucketResponse Response Object
type ShowAppWarehouseBucketResponse struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 是否是用户自建桶
	IsUserCreateBucket *bool `json:"is_user_create_bucket,omitempty"`

	// 桶记录更新时间。
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowAppWarehouseBucketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppWarehouseBucketResponse struct{}"
	}

	return strings.Join([]string{"ShowAppWarehouseBucketResponse", string(data)}, " ")
}
