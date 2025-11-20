package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTransferAssetJobsRequest Request Object
type ListTransferAssetJobsRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 角色。 SENDER:发起方，RECEIVER：接收方。ALL全部
	Role *ListTransferAssetJobsRequestRole `json:"role,omitempty"`

	// 任务状态。多个状态使用英文逗号分隔。 - PROCESSING: 处理过程中 - ACCEPT： 接受 - REJECT： 拒绝 - CANCEL：取消 - FAIL: 失败
	State *string `json:"state,omitempty"`

	// 任务类型。默认查询TRANSFER_OUT类型任务。 - TRANSFER_OUT: 资产转出 - TRANSFER_BACK： 资产转回
	TransferType *string `json:"transfer_type,omitempty"`
}

func (o ListTransferAssetJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransferAssetJobsRequest struct{}"
	}

	return strings.Join([]string{"ListTransferAssetJobsRequest", string(data)}, " ")
}

type ListTransferAssetJobsRequestRole struct {
	value string
}

type ListTransferAssetJobsRequestRoleEnum struct {
	ALL      ListTransferAssetJobsRequestRole
	SENDER   ListTransferAssetJobsRequestRole
	RECEIVER ListTransferAssetJobsRequestRole
}

func GetListTransferAssetJobsRequestRoleEnum() ListTransferAssetJobsRequestRoleEnum {
	return ListTransferAssetJobsRequestRoleEnum{
		ALL: ListTransferAssetJobsRequestRole{
			value: "ALL",
		},
		SENDER: ListTransferAssetJobsRequestRole{
			value: "SENDER",
		},
		RECEIVER: ListTransferAssetJobsRequestRole{
			value: "RECEIVER",
		},
	}
}

func (c ListTransferAssetJobsRequestRole) Value() string {
	return c.value
}

func (c ListTransferAssetJobsRequestRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTransferAssetJobsRequestRole) UnmarshalJSON(b []byte) error {
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
