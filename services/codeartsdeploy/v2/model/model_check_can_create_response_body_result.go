package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCanCreateResponseBodyResult 是否具有创建应用的权限结果
type CheckCanCreateResponseBodyResult struct {

	// 是否具有创建应用的权限
	Creatable *bool `json:"creatable,omitempty"`
}

func (o CheckCanCreateResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCanCreateResponseBodyResult struct{}"
	}

	return strings.Join([]string{"CheckCanCreateResponseBodyResult", string(data)}, " ")
}
