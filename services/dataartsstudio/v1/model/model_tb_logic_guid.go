package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TbLogicGuid 表发布后，创建的数据目录业务资产guid，只读，创建和更新时无需填写。
type TbLogicGuid struct {
}

func (o TbLogicGuid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TbLogicGuid struct{}"
	}

	return strings.Join([]string{"TbLogicGuid", string(data)}, " ")
}
