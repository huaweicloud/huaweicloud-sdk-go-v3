package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseAndDefectInfoResponse Response Object
type ShowTestCaseAndDefectInfoResponse struct {

	// 起始记录数大于实际总条数时，值为0。
	Total *int32 `json:"total,omitempty"`

	// 查询用户用例关联缺陷的统计信息
	Values         *[]ExternalUserCaseAndDefect `json:"values,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowTestCaseAndDefectInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseAndDefectInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowTestCaseAndDefectInfoResponse", string(data)}, " ")
}
