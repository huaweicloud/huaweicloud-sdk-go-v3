package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckFactLogicTableStatusRequest Request Object
type CheckFactLogicTableStatusRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o CheckFactLogicTableStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckFactLogicTableStatusRequest struct{}"
	}

	return strings.Join([]string{"CheckFactLogicTableStatusRequest", string(data)}, " ")
}
