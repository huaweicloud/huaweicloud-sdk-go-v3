package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Page 诊断记录分页查询对象
type Page struct {

	// 一共有多少条数据
	Total *int32 `json:"total,omitempty"`
}

func (o Page) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Page struct{}"
	}

	return strings.Join([]string{"Page", string(data)}, " ")
}
