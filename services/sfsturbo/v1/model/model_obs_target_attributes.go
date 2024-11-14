package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsTargetAttributes 后端存储属性
type ObsTargetAttributes struct {

	// 导入的文件权限。取值范围是0到777。  第一位表示文件所有者的权限，取值范围是0到7；第二位表示文件所属用户组的权限，取值范围是0到7；第三位表示其他用户的权限，取值范围是0到7。文件所有者由uid指定，文件所属用户组由gid指定，不是文件所有者且不在文件所属用户组的用户为其他用户。  数字4、2、1分别表示读、写、执行权限，这些数字相加，即可得到所需的权限组合。例如：750中第一位7代表该文件所有者对该文件具有读、写、执行权限；第二位5代表该文件所属用户组对该文件具有读、执行权限；第三位0代表其他用户对该文件无权限。
	FileMode *int32 `json:"file_mode,omitempty"`

	// 导入的目录权限。取值范围是0到777。  第一位表示目录所有者的权限，取值范围是0到7；第二位表示目录所属用户组的权限，取值范围是0到7；第三位表示其他用户的权限，取值范围是0到7。目录所有者由uid指定，目录所属用户组由gid指定，不是目录所有者且不在目录所属用户组的用户为其他用户。  数字4、2、1分别表示读、写、执行权限，这些数字相加，即可得到所需的权限组合。例如：750中第一位7代表该目录所有者对该目录具有读、写、执行权限；第二位5代表该目录所属用户组对该目录具有读、执行权限；第三位0代表其他用户对该文件无权限。
	DirMode *int32 `json:"dir_mode,omitempty"`

	// 导入对象所有者的用户id，默认值是0，取值范围是0到4,294,967,294（即2^32-2）。
	Uid *int32 `json:"uid,omitempty"`

	// 导入对象所属用户组id，默认值是0，取值范围是0到4,294,967,294（即2^32-2）。
	Gid *int32 `json:"gid,omitempty"`
}

func (o ObsTargetAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsTargetAttributes struct{}"
	}

	return strings.Join([]string{"ObsTargetAttributes", string(data)}, " ")
}
