package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScanPeriodDate 扫描周期日期（1-28；扫描周期为week时，1-7代表周日至周六；扫描周期为month时，1-28代表每月1日到28日）
type ScanPeriodDate struct {
}

func (o ScanPeriodDate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanPeriodDate struct{}"
	}

	return strings.Join([]string{"ScanPeriodDate", string(data)}, " ")
}
