package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BizInfoVo struct {

	// 业务ID，填写String类型替代Long类型。
	BizId string `json:"biz_id"`

	BizType *BizTypeEnum `json:"biz_type"`

	EnvType *EnvTypeEnum `json:"env_type,omitempty"`
}

func (o BizInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizInfoVo struct{}"
	}

	return strings.Join([]string{"BizInfoVo", string(data)}, " ")
}
