package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadGroupPackageReq 上传jar类型分组资源请求body体。
type UploadGroupPackageReq struct {

	// 用户OBS对象路径列表，OBS对象路径为OBS对象URL。
	Paths []string `json:"paths"`

	// 所属资源分组名。
	Group string `json:"group"`
}

func (o UploadGroupPackageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadGroupPackageReq struct{}"
	}

	return strings.Join([]string{"UploadGroupPackageReq", string(data)}, " ")
}
