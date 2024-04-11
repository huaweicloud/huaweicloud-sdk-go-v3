package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogLevelVoList struct {

	// 主题层级信息。
	Levels *[]CatalogLevelVo `json:"levels,omitempty"`
}

func (o CatalogLevelVoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogLevelVoList struct{}"
	}

	return strings.Join([]string{"CatalogLevelVoList", string(data)}, " ")
}
