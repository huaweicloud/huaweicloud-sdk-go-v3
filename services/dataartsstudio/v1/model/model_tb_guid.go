package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TbGuid 表发布后，创建的数据目录技术资产guid，只读，创建和更新时无需填写。
type TbGuid struct {
}

func (o TbGuid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TbGuid struct{}"
	}

	return strings.Join([]string{"TbGuid", string(data)}, " ")
}
