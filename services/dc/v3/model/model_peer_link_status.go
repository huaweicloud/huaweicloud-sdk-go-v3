package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PeerLinkStatus 关联连接状态， - PENDING_CREATE (创建中) - PENDING_UPDATE (更新中) - ACTIVE (可用状态) - ERROR (失败状态)
type PeerLinkStatus struct {
	value string
}

type PeerLinkStatusEnum struct {
	PENDING_CREATE PeerLinkStatus
	PENDING_UPDATE PeerLinkStatus
	ACTIVE         PeerLinkStatus
	ERROR          PeerLinkStatus
}

func GetPeerLinkStatusEnum() PeerLinkStatusEnum {
	return PeerLinkStatusEnum{
		PENDING_CREATE: PeerLinkStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: PeerLinkStatus{
			value: "PENDING_UPDATE",
		},
		ACTIVE: PeerLinkStatus{
			value: "ACTIVE",
		},
		ERROR: PeerLinkStatus{
			value: "ERROR",
		},
	}
}

func (c PeerLinkStatus) Value() string {
	return c.value
}

func (c PeerLinkStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeerLinkStatus) UnmarshalJSON(b []byte) error {
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
