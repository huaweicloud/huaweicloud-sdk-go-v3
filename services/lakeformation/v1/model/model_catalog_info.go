package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogInfo catalog input when grant policy
type CatalogInfo struct {

	// 子数据库信息
	Databases *[]DatabaseInfo `json:"databases,omitempty"`

	// catalog名
	Name string `json:"name"`
}

func (o CatalogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogInfo struct{}"
	}

	return strings.Join([]string{"CatalogInfo", string(data)}, " ")
}
