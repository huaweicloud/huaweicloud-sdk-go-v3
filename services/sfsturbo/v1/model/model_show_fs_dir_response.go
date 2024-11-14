package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFsDirResponse Response Object
type ShowFsDirResponse struct {

	// 目录全路径
	Path *string `json:"path,omitempty"`

	// 目录权限，仅20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/s/TiB、HPC缓存型文件系统返回该字段。第三位表示目录所有者的权限，第四位表示目录所属用户组的权限，第五位表示其他用户的权限。目录所有者由uid指定，目录所属用户组由gid指定，不是目录所有者且不在目录所属用户组的用户为其他用户。例如：40755中第三位7代表目录所有者对该目录具有读、写、执行权限；第四位5代表目录所属用户组对该目录具有读、执行权限；第五位5代表其他用户对该目录具有读、执行权限。
	Mode *int64 `json:"mode,omitempty"`

	// 目录所有者的用户id，仅20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/s/TiB、HPC缓存型文件系统返回该字段。
	Uid *int64 `json:"uid,omitempty"`

	// 目录所属用户组id，仅20MB/s/TiB、40MB/s/TiB、125MB/s/TiB、250MB/s/TiB、500MB/s/TiB、1000MB/s/TiB、HPC缓存型文件系统返回该字段。
	Gid            *int64 `json:"gid,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowFsDirResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsDirResponse struct{}"
	}

	return strings.Join([]string{"ShowFsDirResponse", string(data)}, " ")
}
