package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOmUrlResponse Response Object
type ShowOmUrlResponse struct {

	// 链接获取状态     # SUCCESS(0): 成功；非SUCCESS状态都认为获取失败，需要展示错误代码及description     # IAM_USER_CONFLICT(1016): IAM用户冲突；     # HOST_NOT_MANAGE(1): 查询主机未被纳管；     # HOST_ACCOUNT_NOT_EXIST(553): 主机账户不可用；     # IAM_USER_NO_PERMISSION(901): IAM用户无运维该主机权限；     # CUR_VERSION_NOT_SUPPORT_OPERATION(9001):当前服务版本不支持;     # INS_WHITE_LIST_INITIALIZING(9002):实例白名单正在初始化，请稍后重试;     # UNKNOWN_ERROR(9003):未知错误;
	State *ShowOmUrlResponseState `json:"state,omitempty"`

	// 链接获取异常时说明提示
	Description *string `json:"description,omitempty"`

	// ECS运维登录地址
	LoginUrl       *string `json:"login_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowOmUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOmUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowOmUrlResponse", string(data)}, " ")
}

type ShowOmUrlResponseState struct {
	value string
}

type ShowOmUrlResponseStateEnum struct {
	SUCCESS                           ShowOmUrlResponseState
	IAM_USER_CONFLICT                 ShowOmUrlResponseState
	HOST_NOT_MANAGE                   ShowOmUrlResponseState
	HOST_ACCOUNT_NOT_EXIST            ShowOmUrlResponseState
	IAM_USER_NO_PERMISSION            ShowOmUrlResponseState
	CUR_VERSION_NOT_SUPPORT_OPERATION ShowOmUrlResponseState
	INS_WHITE_LIST_INITIALIZING       ShowOmUrlResponseState
	UNKNOWN_ERROR                     ShowOmUrlResponseState
}

func GetShowOmUrlResponseStateEnum() ShowOmUrlResponseStateEnum {
	return ShowOmUrlResponseStateEnum{
		SUCCESS: ShowOmUrlResponseState{
			value: "SUCCESS",
		},
		IAM_USER_CONFLICT: ShowOmUrlResponseState{
			value: "IAM_USER_CONFLICT",
		},
		HOST_NOT_MANAGE: ShowOmUrlResponseState{
			value: "HOST_NOT_MANAGE",
		},
		HOST_ACCOUNT_NOT_EXIST: ShowOmUrlResponseState{
			value: "HOST_ACCOUNT_NOT_EXIST",
		},
		IAM_USER_NO_PERMISSION: ShowOmUrlResponseState{
			value: "IAM_USER_NO_PERMISSION",
		},
		CUR_VERSION_NOT_SUPPORT_OPERATION: ShowOmUrlResponseState{
			value: "CUR_VERSION_NOT_SUPPORT_OPERATION",
		},
		INS_WHITE_LIST_INITIALIZING: ShowOmUrlResponseState{
			value: "INS_WHITE_LIST_INITIALIZING",
		},
		UNKNOWN_ERROR: ShowOmUrlResponseState{
			value: "UNKNOWN_ERROR",
		},
	}
}

func (c ShowOmUrlResponseState) Value() string {
	return c.value
}

func (c ShowOmUrlResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOmUrlResponseState) UnmarshalJSON(b []byte) error {
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
