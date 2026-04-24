package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCompareUsersDetailResponse Response Object
type ShowCompareUsersDetailResponse struct {

	// 用户对比信息的总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 用户对比信息
	UserCompareInfo *[]CompareUserInfo `json:"user_compare_info,omitempty"`
	HttpStatusCode  int                `json:"-"`
}

func (o ShowCompareUsersDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompareUsersDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCompareUsersDetailResponse", string(data)}, " ")
}
