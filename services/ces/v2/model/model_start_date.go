package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDate 屏蔽起始日期，yyyy-MM-dd。
type StartDate struct {
}

func (o StartDate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDate struct{}"
	}

	return strings.Join([]string{"StartDate", string(data)}, " ")
}
