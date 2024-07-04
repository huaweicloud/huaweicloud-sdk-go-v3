package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserSyncReq struct {

	// StarRocks账号同步动作。
	Action UserSyncReqAction `json:"action"`
}

func (o UserSyncReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserSyncReq struct{}"
	}

	return strings.Join([]string{"UserSyncReq", string(data)}, " ")
}

type UserSyncReqAction struct {
	value string
}

type UserSyncReqActionEnum struct {
	START_SYNC_TAURUS_ACCOUNT UserSyncReqAction
	STOP_SYNC_TAURUS_ACCOUNT  UserSyncReqAction
}

func GetUserSyncReqActionEnum() UserSyncReqActionEnum {
	return UserSyncReqActionEnum{
		START_SYNC_TAURUS_ACCOUNT: UserSyncReqAction{
			value: "startSyncTaurusAccount",
		},
		STOP_SYNC_TAURUS_ACCOUNT: UserSyncReqAction{
			value: "stopSyncTaurusAccount",
		},
	}
}

func (c UserSyncReqAction) Value() string {
	return c.value
}

func (c UserSyncReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserSyncReqAction) UnmarshalJSON(b []byte) error {
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
