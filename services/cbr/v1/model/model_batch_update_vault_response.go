package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateVaultResponse Response Object
type BatchUpdateVaultResponse struct {

	// 已批量修改id列表
	UpdatedVaultsId *[]string `json:"updated_vaults_id,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o BatchUpdateVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateVaultResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateVaultResponse", string(data)}, " ")
}
