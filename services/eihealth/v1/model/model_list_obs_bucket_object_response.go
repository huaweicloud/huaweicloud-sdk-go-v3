package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListObsBucketObjectResponse struct {

	// 数据（文件夹、文件）总数量
	Count *int32 `json:"count,omitempty"`

	// 数据列表
	Objects        *[]BucketObjectDto `json:"objects,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListObsBucketObjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketObjectResponse struct{}"
	}

	return strings.Join([]string{"ListObsBucketObjectResponse", string(data)}, " ")
}
