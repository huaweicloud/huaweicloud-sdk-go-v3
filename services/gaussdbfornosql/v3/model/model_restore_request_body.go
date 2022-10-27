package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 恢复到已有实例的请求body
type RestoreRequestBody struct {

	// 备份文件名称。根据备份文件恢复到已有的实例。
	BackupId string `json:"backup_id"`

	// 实例密码。 取值范围：长度为8~32位，必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。   - 不传入密码时，恢复后，备份文件中保留的密码将覆盖原有实例的密码。   - 传入密码时，恢复后，将使用该密码覆盖原有实例的密码。
	Password *string `json:"password,omitempty"`
}

func (o RestoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreRequestBody", string(data)}, " ")
}
