package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutoSnapshotPolicyResponse Response Object
type DeleteAutoSnapshotPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAutoSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAutoSnapshotPolicyResponse", string(data)}, " ")
}
