package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyV5Request Request Object
type DeleteAgencyV5Request struct {

	// 信任委托ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	AgencyId string `json:"agency_id"`
}

func (o DeleteAgencyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyV5Request struct{}"
	}

	return strings.Join([]string{"DeleteAgencyV5Request", string(data)}, " ")
}
