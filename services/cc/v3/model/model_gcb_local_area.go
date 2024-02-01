package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GcbLocalArea 全域互联带宽本端接入点。
type GcbLocalArea struct {

	// 功能说明：本端接入点，配合remote_area信息描述带宽实例应用的范围。 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点），站点编码通过接口获取，带宽类型为Region可不传，其他类型必传
	LocalArea *string `json:"local_area,omitempty"`
}

func (o GcbLocalArea) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbLocalArea struct{}"
	}

	return strings.Join([]string{"GcbLocalArea", string(data)}, " ")
}
