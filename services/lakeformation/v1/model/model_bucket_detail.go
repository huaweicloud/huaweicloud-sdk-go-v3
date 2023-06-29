package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BucketDetail obs桶详细信息
type BucketDetail struct {

	// obs桶id
	Id string `json:"id"`

	// obs桶名
	Name string `json:"name"`

	// obs桶创建时间
	CreationDate *sdktime.SdkTime `json:"creation_date,omitempty"`

	// 所属region
	Location string `json:"location"`
}

func (o BucketDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketDetail struct{}"
	}

	return strings.Join([]string{"BucketDetail", string(data)}, " ")
}
