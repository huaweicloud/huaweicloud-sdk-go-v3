package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteGroupReferenceInfo 站点分组跟外部关联的数据模型。
type SiteGroupReferenceInfo struct {

	// 资源ID标识符。
	Id string `json:"id"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`
}

func (o SiteGroupReferenceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteGroupReferenceInfo struct{}"
	}

	return strings.Join([]string{"SiteGroupReferenceInfo", string(data)}, " ")
}
