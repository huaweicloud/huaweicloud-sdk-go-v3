package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalEip 创建全域弹性公网IP响应体
type CreateGlobalEip struct {

	// 全域弹性公网IP的ID
	Id string `json:"id"`

	// - 功能说明：全域弹性公网IP段名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name string `json:"name"`
}

func (o CreateGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEip struct{}"
	}

	return strings.Join([]string{"CreateGlobalEip", string(data)}, " ")
}
