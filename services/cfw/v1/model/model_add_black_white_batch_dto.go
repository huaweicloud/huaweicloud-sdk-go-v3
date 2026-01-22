package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddBlackWhiteBatchDto struct {

	// **参数解释**： 黑白名单信息列表 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	ListItems []BlackWhiteInfo `json:"list_items"`

	// **参数解释**： 黑白名单列表类型 **约束限制**： 不涉及 **取值范围**： 4：黑名单 5：白名单 **默认取值**： 不涉及
	ListType int32 `json:"list_type"`

	// **参数解释**： 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	ObjectId string `json:"object_id"`
}

func (o AddBlackWhiteBatchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteBatchDto struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteBatchDto", string(data)}, " ")
}
