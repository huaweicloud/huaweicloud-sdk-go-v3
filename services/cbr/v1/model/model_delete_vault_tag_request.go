package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteVaultTagRequest struct {
	// 不能为空或空字符串，不检查长度和字符集，去掉key前后的空格后检查，去掉key前后的空格后使用。 即使底层存在非法的tag也要能删。

	Key string `json:"key"`
	// 资源id

	VaultId string `json:"vault_id"`
}

func (o DeleteVaultTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteVaultTagRequest", string(data)}, " ")
}
