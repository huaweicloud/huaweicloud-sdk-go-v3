package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsBucketObjectsResponse Response Object
type ListObsBucketObjectsResponse struct {

	// 对象的总数
	Count *int32 `json:"count,omitempty"`

	// obs文件
	Objects *[]ObsObject `json:"objects,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListObsBucketObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListObsBucketObjectsResponse", string(data)}, " ")
}
