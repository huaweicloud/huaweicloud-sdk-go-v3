package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasResource 资源配额详情。
type QuotasResource struct {

	// 项目资源类型。
	Type string `json:"type"`

	// 已使用的资源数量。
	Used string `json:"used"`

	// 项目资源配额。
	Quota int32 `json:"quota"`

	// 资源计量单位。
	Unit int32 `json:"unit"`
}

func (o QuotasResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasResource struct{}"
	}

	return strings.Join([]string{"QuotasResource", string(data)}, " ")
}
