package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteGroupReferenceInfo 站点分组跟外部关联的数据模型。
type SiteGroupReferenceInfo struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 功能说明：站点分组自定义的英文名字。 取值范围：1-255个字符。
	NameEn *string `json:"name_en,omitempty"`

	// 功能说明：站点分组自定义的中文名字。 取值范围：1-64个字符。
	NameCn *string `json:"name_cn,omitempty"`
}

func (o SiteGroupReferenceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteGroupReferenceInfo struct{}"
	}

	return strings.Join([]string{"SiteGroupReferenceInfo", string(data)}, " ")
}
