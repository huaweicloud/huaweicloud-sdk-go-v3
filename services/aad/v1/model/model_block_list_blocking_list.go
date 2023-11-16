package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BlockListBlockingList struct {

	// ip地址
	Ip string `json:"ip"`

	// 封堵时间
	BlockingTime int64 `json:"blocking_time"`

	// 预计解封时间
	EstimatedUnblockingTime int64 `json:"estimated_unblocking_time"`

	// 状态。unblocking：解封中；success：成功；failed：失败
	Status BlockListBlockingListStatus `json:"status"`
}

func (o BlockListBlockingList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlockListBlockingList struct{}"
	}

	return strings.Join([]string{"BlockListBlockingList", string(data)}, " ")
}

type BlockListBlockingListStatus struct {
	value string
}

type BlockListBlockingListStatusEnum struct {
	FAILED     BlockListBlockingListStatus
	SUCCESS    BlockListBlockingListStatus
	UNBLOCKING BlockListBlockingListStatus
}

func GetBlockListBlockingListStatusEnum() BlockListBlockingListStatusEnum {
	return BlockListBlockingListStatusEnum{
		FAILED: BlockListBlockingListStatus{
			value: "failed",
		},
		SUCCESS: BlockListBlockingListStatus{
			value: "success",
		},
		UNBLOCKING: BlockListBlockingListStatus{
			value: "unblocking",
		},
	}
}

func (c BlockListBlockingListStatus) Value() string {
	return c.value
}

func (c BlockListBlockingListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BlockListBlockingListStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
