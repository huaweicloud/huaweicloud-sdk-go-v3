package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSkillPackage 创建企业自研技能时的技能包信息。
type CreateSkillPackage struct {

	// 版本号。
	Version string `json:"version"`

	// 技能包文件名（与 getUploadUrls 中的 packageName 一致），服务端据此构造 OBS 路径。
	PackageName string `json:"package_name"`

	// 技能包 SHA256 哈希值（前端上传前计算）。
	PackageHash string `json:"package_hash"`

	// 技能包大小（字节），最大 104857600。
	PackageSize int64 `json:"package_size"`

	// OBS 存储区域信息（含上传状态）。
	Regions []PackageRegionWithStatusInfo `json:"regions"`
}

func (o CreateSkillPackage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillPackage struct{}"
	}

	return strings.Join([]string{"CreateSkillPackage", string(data)}, " ")
}
