package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSpecifiedVersionDetailsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type"`

	// 待查询版本号。取值 以v开头，例如v1。 若为空，表示查询所 有API的版本号。
	Version string `json:"version"`
}

func (o ListSpecifiedVersionDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecifiedVersionDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSpecifiedVersionDetailsRequest", string(data)}, " ")
}
