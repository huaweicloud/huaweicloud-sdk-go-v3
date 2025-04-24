package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddTestCaseResultLogInfo struct {

	// 版本号
	ReleaseDev *string `json:"release_dev,omitempty"`
}

func (o AddTestCaseResultLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTestCaseResultLogInfo struct{}"
	}

	return strings.Join([]string{"AddTestCaseResultLogInfo", string(data)}, " ")
}
