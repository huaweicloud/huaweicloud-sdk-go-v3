package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InnodbLock struct {

	// 锁ID
	LockId string `json:"lock_id" xml:"lock_id"`

	// 事务ID
	LockTrxId string `json:"lock_trx_id" xml:"lock_trx_id"`

	// 锁模式，取值为S[,GAP], X[,GAP], IS[,GAP], IX[,GAP], AUTO_INC, and UNKNOWN。
	LockMode string `json:"lock_mode" xml:"lock_mode"`

	// 锁类型，取值为RECORD或TABLE。RECORD为行锁, TABLE为表锁
	LockType string `json:"lock_type" xml:"lock_type"`

	// 加锁的表
	LockTable string `json:"lock_table" xml:"lock_table"`

	// 如果是lock_type='RECORD' 行级锁 ,为锁住的索引，如果是表锁为null
	LockIndex string `json:"lock_index" xml:"lock_index"`

	// 如果是lock_type='RECORD' 行级锁 ,为锁住的索引，如果是表锁为null
	LockSpace string `json:"lock_space" xml:"lock_space"`

	// 如果是lock_type='RECORD' 行级锁 ,为锁住的页号，如果是表锁为null
	LockPage string `json:"lock_page" xml:"lock_page"`

	// 如果是lock_type='RECORD' 行级锁 ,为锁住的堆号，如果是表锁为null
	LockRec string `json:"lock_rec" xml:"lock_rec"`

	// 事务锁住的主键值，若是表锁，则该值为null
	LockData string `json:"lock_data" xml:"lock_data"`
}

func (o InnodbLock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InnodbLock struct{}"
	}

	return strings.Join([]string{"InnodbLock", string(data)}, " ")
}
