package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// permission列表。
type Permission struct {

	// permission的ID，唯一标识。
	Id *string `json:"id,omitempty" xml:"id"`

	// permission列表。 权限格式为“iam:domain:: 6e9dfd51d1124e8d8498dce894923a0d”或 “*”，“*”表示所有用户的终端节点可连接。 其中6e9dfd51d1124e8d8498dce894923a0d为 可连接的用户domian_id。
	Permission *string `json:"permission,omitempty" xml:"permission"`

	// 白名单的添加时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH: MM:SSZ
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`
}

func (o Permission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Permission struct{}"
	}

	return strings.Join([]string{"Permission", string(data)}, " ")
}
