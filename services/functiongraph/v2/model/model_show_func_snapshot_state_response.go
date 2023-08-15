package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFuncSnapshotStateResponse Response Object
type ShowFuncSnapshotStateResponse struct {

	// 快照制作状态
	State *string `json:"state,omitempty"`

	// 快照制作响应码
	Code           *string `json:"code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFuncSnapshotStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFuncSnapshotStateResponse struct{}"
	}

	return strings.Join([]string{"ShowFuncSnapshotStateResponse", string(data)}, " ")
}
