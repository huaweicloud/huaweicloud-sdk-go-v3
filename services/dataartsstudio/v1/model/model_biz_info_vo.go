package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BizInfoVo struct {

	// 业务ID
	BizId int64 `json:"biz_id"`

	BizType *BizTypeEnum `json:"biz_type"`
}

func (o BizInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizInfoVo struct{}"
	}

	return strings.Join([]string{"BizInfoVo", string(data)}, " ")
}
