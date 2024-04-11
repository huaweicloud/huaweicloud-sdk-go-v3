package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckFactLogicTableStatusResponse Response Object
type CheckFactLogicTableStatusResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CheckFactLogicTableStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckFactLogicTableStatusResponse struct{}"
	}

	return strings.Join([]string{"CheckFactLogicTableStatusResponse", string(data)}, " ")
}
