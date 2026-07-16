package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceIdName 服务id和name信息
type ServiceIdName struct {

	// **参数解释：** 服务ID。 **取值范围：** 服务ID。
	Id string `json:"id"`

	// **参数解释：** 服务名。 **取值范围：** 支持1-128个字符，可以包含字母、汉字、数字、连字符和下划线。
	Name string `json:"name"`
}

func (o ServiceIdName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceIdName struct{}"
	}

	return strings.Join([]string{"ServiceIdName", string(data)}, " ")
}
