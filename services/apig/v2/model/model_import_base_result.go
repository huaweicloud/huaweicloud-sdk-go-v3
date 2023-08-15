package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportBaseResult struct {

	// 导入成功信息
	Success *[]Success `json:"success,omitempty"`

	// 导入失败信息
	Failure *[]Failure `json:"failure,omitempty"`

	Swagger *Swagger `json:"swagger,omitempty"`
}

func (o ImportBaseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBaseResult struct{}"
	}

	return strings.Join([]string{"ImportBaseResult", string(data)}, " ")
}
