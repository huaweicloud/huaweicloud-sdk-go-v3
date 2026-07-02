package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FromTimeStampSchema **参数解释** 查询数据起始时间，UNIX时间戳，单位毫秒 **约束限制** 当period为1时，若(to- from) >4*3600*1000，则from调整为 to - 4*3600*1000 当period为300时，若(to - from) >24*3600*1000，则from调整为 to - 24*3600*1000 当period为1200时，若(to - from) >3*24*3600*1000，则from调整为 to - 3*24*3600*1000 当period为3600时，若(to -from) > 10*24*3600*1000，则from调整为 to -10*24*3600*1000 当period为14400时，若(to - from) >30*24*3600*1000，则from调整为 to - 30*24*3600*1000 当period为86400时，若(to -from) > 180*24*3600*1000，则from调整为 to - 180*24*3600*1000 **取值范围** 毫秒级时间戳范围为[1111111111111,9999999999999] **默认取值** 不涉及
type FromTimeStampSchema struct {
}

func (o FromTimeStampSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FromTimeStampSchema struct{}"
	}

	return strings.Join([]string{"FromTimeStampSchema", string(data)}, " ")
}
