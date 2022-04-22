package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 场地条件
type Condition struct {

	// 机房环境条件 取值范围：   - 0：需做详细评估   - 1：机房内已部署华为FusionModule   - 2：机房等级满足T3或同等级国家标准
	Environment *int32 `json:"environment,omitempty"`

	// 机柜空间条件 取值范围：   - 0：除首柜外，同机房无预留空间   - 1：除首柜外，同机房预留3柜以上空间   - 2：除首柜外，同机房预留1-3柜空间
	Space *int32 `json:"space,omitempty"`

	// 运输条件 取值范围：   - 0：运输通道和机房门的高度或宽度不满足要求   - 1：运输通道，货梯，机房门均可满足整机柜滚轮搬运   - 2：运输通道，货梯，机房门不能支持整机柜滚轮搬运，沿途有台阶
	Transport *int32 `json:"transport,omitempty"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}
