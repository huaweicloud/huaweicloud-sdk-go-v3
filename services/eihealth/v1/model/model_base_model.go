package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseModel 基模型
type BaseModel struct {

	// 基模型id
	Id string `json:"id"`

	// 基模型名称
	Name string `json:"name"`
}

func (o BaseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseModel struct{}"
	}

	return strings.Join([]string{"BaseModel", string(data)}, " ")
}
