package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUploadUrlsReq 获取 OBS 预签名上传地址请求。
type CreateUploadUrlsReq struct {

	// 技能slug（用于构建 OBS 路径）。
	Slug string `json:"slug"`

	// 版本号（用于构建 OBS 路径）。
	Version string `json:"version"`

	// 包文件名。
	PackageName string `json:"package_name"`

	// 目标 region 列表。
	Regions []string `json:"regions"`
}

func (o CreateUploadUrlsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUploadUrlsReq struct{}"
	}

	return strings.Join([]string{"CreateUploadUrlsReq", string(data)}, " ")
}
