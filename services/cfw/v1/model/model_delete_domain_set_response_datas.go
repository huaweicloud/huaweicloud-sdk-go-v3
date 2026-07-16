package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDomainSetResponseDatas struct {

	// **参数解释**： 批量删除域名组响应信息 **取值范围**： 不涉及
	ResponseDatas *[]DomainSetId `json:"responseDatas,omitempty"`
}

func (o DeleteDomainSetResponseDatas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainSetResponseDatas struct{}"
	}

	return strings.Join([]string{"DeleteDomainSetResponseDatas", string(data)}, " ")
}
