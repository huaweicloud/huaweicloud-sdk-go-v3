package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddOrModifyTagReq struct {

	// 标签名称
	TagName string `json:"tag_name"`
}

func (o AddOrModifyTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrModifyTagReq struct{}"
	}

	return strings.Join([]string{"AddOrModifyTagReq", string(data)}, " ")
}
