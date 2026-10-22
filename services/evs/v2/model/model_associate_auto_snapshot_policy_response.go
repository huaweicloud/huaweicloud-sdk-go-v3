package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAutoSnapshotPolicyResponse Response Object
type AssociateAutoSnapshotPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateAutoSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAutoSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateAutoSnapshotPolicyResponse", string(data)}, " ")
}
