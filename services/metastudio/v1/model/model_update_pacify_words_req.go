package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePacifyWordsReq 修改安抚话术请求。
type UpdatePacifyWordsReq struct {

	// 安抚话术。
	PacifyWords string `json:"pacify_words"`
}

func (o UpdatePacifyWordsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePacifyWordsReq struct{}"
	}

	return strings.Join([]string{"UpdatePacifyWordsReq", string(data)}, " ")
}
