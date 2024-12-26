package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResponseData 通用返回体
type ResponseData struct {

	// 数据id
	Id *string `json:"id,omitempty"`
}

func (o ResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseData struct{}"
	}

	return strings.Join([]string{"ResponseData", string(data)}, " ")
}
