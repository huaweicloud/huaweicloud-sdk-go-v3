package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaselineCatalogModel struct {

	// 目录ID唯一标识，UUID
	Uuid string `json:"uuid"`

	// 目录的位置顺序
	SerialNumber int32 `json:"serial_number"`

	// 目录的层级关系
	LevelNumber int32 `json:"level_number"`

	// 该目录所在遵从包UUID
	Root string `json:"root"`

	// 该目录的父目录UUID，如果等于为第一层目录，则为遵从包UUID
	Parent string `json:"parent"`

	// 该目录是否是叶子节点 0：不是 1：是
	IsLeaf *bool `json:"is_leaf,omitempty"`

	// 目录关联的检查项
	CheckItems *[]CheckitemCatalogModel `json:"check_items,omitempty"`
}

func (o BaselineCatalogModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaselineCatalogModel struct{}"
	}

	return strings.Join([]string{"BaselineCatalogModel", string(data)}, " ")
}
