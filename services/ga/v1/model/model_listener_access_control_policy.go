package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListenerAccessControlPolicy 监听器的访问控制策略。
type ListenerAccessControlPolicy struct {

	// 监听器ID。
	ListenerId *string `json:"listener_id,omitempty"`

	Type *ListenerAccessControlType `json:"type,omitempty"`
}

func (o ListenerAccessControlPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerAccessControlPolicy struct{}"
	}

	return strings.Join([]string{"ListenerAccessControlPolicy", string(data)}, " ")
}
