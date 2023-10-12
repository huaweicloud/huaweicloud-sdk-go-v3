package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndDate 屏蔽截止日期，yyyy-MM-dd。
type EndDate struct {
}

func (o EndDate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndDate struct{}"
	}

	return strings.Join([]string{"EndDate", string(data)}, " ")
}
