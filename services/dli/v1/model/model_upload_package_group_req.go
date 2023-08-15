package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPackageGroupReq 上传分组资源的请求body体。
type UploadPackageGroupReq struct {

	// 用户OBS对象路径列表，OBS对象路径为OBS对象URL。
	Paths []string `json:"paths"`

	// 分组资源文件的类型。 说明：上传的同一组资源包含不同文件类型时，均选择“file”类型作为这次上传文件的类型。
	Kind string `json:"kind"`

	// 将要创建的分组名。
	Group string `json:"group"`

	// 是否异步上传资源包
	IsAsync *bool `json:"is_async,omitempty"`

	// 资源标签。具体请参考表tags。
	Tags *[]TmsTagEntity `json:"tags,omitempty"`
}

func (o UploadPackageGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPackageGroupReq struct{}"
	}

	return strings.Join([]string{"UploadPackageGroupReq", string(data)}, " ")
}
