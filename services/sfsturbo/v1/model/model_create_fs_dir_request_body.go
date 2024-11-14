package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFsDirRequestBody 创建目录请求
type CreateFsDirRequestBody struct {

	// 合法的的目录全路径
	Path string `json:"path"`

	// 目录权限，默认值是755，取值范围是0到777。第一位表示目录所有者的权限，取值范围是0到7；第二位表示目录所属用户组的权限，取值范围是0到7；第三位表示其他用户的权限，取值范围是0到7。目录所有者由uid指定，目录所属用户组由gid指定，不是目录所有者且不在目录所属用户组的用户为其他用户。例如：755中第一位7代表该目录所有者对该目录具有读、写、执行权限；第二位5代表该目录所属用户组对该目录具有读、执行权限；第三位5代表其他用户对该目录具有读、执行权限。
	Mode *int64 `json:"mode,omitempty"`

	// 目录所有者的用户id，默认值是0，取值范围是0到4,294,967,294（即2^32-2）。
	Uid *int64 `json:"uid,omitempty"`

	// 目录所属用户组id，默认值是0，取值范围是0到4,294,967,294（即2^32-2）。
	Gid *int64 `json:"gid,omitempty"`
}

func (o CreateFsDirRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsDirRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFsDirRequestBody", string(data)}, " ")
}
