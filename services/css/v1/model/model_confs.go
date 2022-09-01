package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Confs struct {

	// 配置文件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 配置文件状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 配置文件内容。
	ConfContent *string `json:"confContent,omitempty" xml:"confContent"`

	Setting *Setting `json:"setting,omitempty" xml:"setting"`

	// 更新时间。
	UpdateAt *string `json:"updateAt,omitempty" xml:"updateAt"`
}

func (o Confs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Confs struct{}"
	}

	return strings.Join([]string{"Confs", string(data)}, " ")
}
