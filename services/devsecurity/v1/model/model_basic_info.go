package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicInfo struct {

	// 应用名称，移动应用特有
	AppName *string `json:"app_name,omitempty"`

	// 文件包名
	PackageName *string `json:"package_name,omitempty"`

	// 文件大小
	FileSize *int32 `json:"file_size,omitempty"`

	// 版本号
	VersionCode *int32 `json:"version_code,omitempty"`

	// 最小SDK版本
	MinSdk *string `json:"min_sdk,omitempty"`

	// 目标SDK版本
	TargetSdk *string `json:"target_sdk,omitempty"`

	// 文件SHA1值
	Sha1 *string `json:"sha1,omitempty"`

	// 文件SHA256值
	Sha256 *string `json:"sha256,omitempty"`

	// 文件MD5值
	Md5 *string `json:"md5,omitempty"`

	// 证书发布者
	Issuer *string `json:"issuer,omitempty"`

	// 证书拥有者
	Owner *string `json:"owner,omitempty"`

	// 证书有效日期
	EffectiveDate *string `json:"effective_date,omitempty"`

	// 算法
	Algorithm *string `json:"algorithm,omitempty"`

	// 证书公钥
	PublicKey *string `json:"public_key,omitempty"`
}

func (o BasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicInfo struct{}"
	}

	return strings.Join([]string{"BasicInfo", string(data)}, " ")
}
