package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSkillPackageReq 上传技能包请求（企业自研技能，替换/升级版本）。
type CreateSkillPackageReq struct {

	// 版本号。
	Version string `json:"version"`

	// 技能包文件名（与 getUploadUrls 中的 packageName 一致），服务端据此构造 OBS 路径。
	PackageName string `json:"package_name"`

	// 技能包 SHA256 哈希值（前端上传前计算）。
	PackageHash string `json:"package_hash"`

	// 技能包大小（字节）。
	PackageSize int64 `json:"package_size"`

	// 区域信息（含上传状态）。
	Regions []PackageRegionWithStatusInfo `json:"regions"`

	// 备注。
	Remark *string `json:"remark,omitempty"`
}

func (o CreateSkillPackageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillPackageReq struct{}"
	}

	return strings.Join([]string{"CreateSkillPackageReq", string(data)}, " ")
}
