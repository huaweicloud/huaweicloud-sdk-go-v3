package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListObsBucketRequest struct {

	// 桶类型。OBJECT：桶列表；PFS：并行文件系统。不传返回所有
	Type *string `json:"type,omitempty"`
}

func (o ListObsBucketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketRequest struct{}"
	}

	return strings.Join([]string{"ListObsBucketRequest", string(data)}, " ")
}
