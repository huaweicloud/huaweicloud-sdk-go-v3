package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSparseBackupPolicyReq struct {

	// 备份策略集
	Policies []SparseBackupPolicyForUpdate `json:"policies"`
}

func (o UpdateSparseBackupPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSparseBackupPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateSparseBackupPolicyReq", string(data)}, " ")
}
