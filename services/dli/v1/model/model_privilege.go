package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Privilege struct {

	// 赋权的DLI资源，具体可参考https://support.huaweicloud.com/usermanual-dli/dli_01_0417.html
	Object string `json:"object"`

	// 待赋权，回收或更新的权限列表。 说明： 若“action”为“update”，更新列表为空，则表示回收用户在DLI的所有资源权限。
	Privileges []string `json:"privileges"`
}

func (o Privilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Privilege struct{}"
	}

	return strings.Join([]string{"Privilege", string(data)}, " ")
}
