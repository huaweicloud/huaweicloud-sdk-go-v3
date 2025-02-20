package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StepImageVo 实际的数据类型：单个对象，集合 或 NULL
type StepImageVo struct {

	// 测试步骤图片id
	ImgId *string `json:"img_id,omitempty"`

	// 测试步骤图片路径
	ImgUrl *string `json:"img_url,omitempty"`
}

func (o StepImageVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepImageVo struct{}"
	}

	return strings.Join([]string{"StepImageVo", string(data)}, " ")
}
