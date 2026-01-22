package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateIpsWhiteListDto struct {

	// **参数解释**：  添加的IPS白名单列表 **约束限制**：  不涉及  **取值范围**： 不涉及  **默认取值**：  不涉及
	IpsWhiteListDtoList []IpsWhiteListDto `json:"ips_white_list_dto_list"`
}

func (o BatchCreateIpsWhiteListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIpsWhiteListDto struct{}"
	}

	return strings.Join([]string{"BatchCreateIpsWhiteListDto", string(data)}, " ")
}
