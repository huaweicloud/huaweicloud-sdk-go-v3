package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDataRestInfoImageInfoObjects struct {

	// 主体目标框。
	Box *string `json:"box,omitempty"`

	// 主体类目序号。
	Category float32 `json:"category,omitempty"`

	// 主体类目名称。
	CategoryName *string `json:"category_name,omitempty"`
}

func (o AddDataRestInfoImageInfoObjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDataRestInfoImageInfoObjects struct{}"
	}

	return strings.Join([]string{"AddDataRestInfoImageInfoObjects", string(data)}, " ")
}
