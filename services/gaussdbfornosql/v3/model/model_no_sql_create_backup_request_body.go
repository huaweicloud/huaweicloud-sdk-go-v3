package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NoSqlCreateBackupRequestBody 创建手动备份请求参数
type NoSqlCreateBackupRequestBody struct {

	// 手动备份名称。  取值范围：长度为4~64位，必须以字母开头（A~Z或a~z），区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。
	Name string `json:"name"`

	// 手动备份描述。  取值范围：长度不超过256位，且不能包含>!<\"&'=特殊字符。
	Description string `json:"description"`
}

func (o NoSqlCreateBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlCreateBackupRequestBody struct{}"
	}

	return strings.Join([]string{"NoSqlCreateBackupRequestBody", string(data)}, " ")
}
