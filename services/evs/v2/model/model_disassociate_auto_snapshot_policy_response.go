package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAutoSnapshotPolicyResponse Response Object
type DisassociateAutoSnapshotPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisassociateAutoSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAutoSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisassociateAutoSnapshotPolicyResponse", string(data)}, " ")
}
