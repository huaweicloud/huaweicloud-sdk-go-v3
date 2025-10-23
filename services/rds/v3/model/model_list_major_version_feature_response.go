package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMajorVersionFeatureResponse Response Object
type ListMajorVersionFeatureResponse struct {

	// 版本名称。
	Version *string `json:"version,omitempty"`

	// 支持升级的版本列表。
	SupportUpgradeVersion *[]string `json:"support_upgrade_version,omitempty"`

	// 支持备份恢复的版本列表。
	SupportRecoverVersion *[]string `json:"support_recover_version,omitempty"`

	// 是否支持FileStream特性。
	SupportFileStream *bool `json:"support_file_stream,omitempty"`

	// 是否支持透明数据加密。
	SupportTde *bool `json:"support_tde,omitempty"`

	// 是否支持Always On。
	SupportAlwaysOn *bool `json:"support_always_on,omitempty"`

	// 是否支持只读。
	SupportReadOnly *bool `json:"support_read_only,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListMajorVersionFeatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMajorVersionFeatureResponse struct{}"
	}

	return strings.Join([]string{"ListMajorVersionFeatureResponse", string(data)}, " ")
}
