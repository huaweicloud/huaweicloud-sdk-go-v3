package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GcbRemoteArea 全域互联带宽远端接入点。
type GcbRemoteArea struct {

	// 功能说明：远端接入点，配合local_area信息描述带宽实例应用的范围。 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点），站点编码通过接口获取，带宽类型为Region可不传，其他类型必传
	RemoteArea *string `json:"remote_area,omitempty"`
}

func (o GcbRemoteArea) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbRemoteArea struct{}"
	}

	return strings.Join([]string{"GcbRemoteArea", string(data)}, " ")
}
