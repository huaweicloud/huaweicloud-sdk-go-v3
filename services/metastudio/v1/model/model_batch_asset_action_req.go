package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchAssetActionReq 批量操作请求体
type BatchAssetActionReq struct {

	// 批量操作命令. * DELETE：删除 * DELETE_FORCE：强制删除，该模式会立即删除资产 * RESTORE：恢复 * UNACTIVE：取消激活 * ACTIVE：激活 * SHARE：共享 * UNSHARE：取消共享
	Action BatchAssetActionReqAction `json:"action"`

	// 资产ID列表
	AssetIds []string `json:"asset_ids"`

	// 操作参数
	Params *string `json:"params,omitempty"`
}

func (o BatchAssetActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssetActionReq struct{}"
	}

	return strings.Join([]string{"BatchAssetActionReq", string(data)}, " ")
}

type BatchAssetActionReqAction struct {
	value string
}

type BatchAssetActionReqActionEnum struct {
	DELETE       BatchAssetActionReqAction
	DELETE_FORCE BatchAssetActionReqAction
	RESTORE      BatchAssetActionReqAction
	UNACTIVE     BatchAssetActionReqAction
	ACTIVE       BatchAssetActionReqAction
	SHARE        BatchAssetActionReqAction
	UNSHARE      BatchAssetActionReqAction
}

func GetBatchAssetActionReqActionEnum() BatchAssetActionReqActionEnum {
	return BatchAssetActionReqActionEnum{
		DELETE: BatchAssetActionReqAction{
			value: "DELETE",
		},
		DELETE_FORCE: BatchAssetActionReqAction{
			value: "DELETE_FORCE",
		},
		RESTORE: BatchAssetActionReqAction{
			value: "RESTORE",
		},
		UNACTIVE: BatchAssetActionReqAction{
			value: "UNACTIVE",
		},
		ACTIVE: BatchAssetActionReqAction{
			value: "ACTIVE",
		},
		SHARE: BatchAssetActionReqAction{
			value: "SHARE",
		},
		UNSHARE: BatchAssetActionReqAction{
			value: "UNSHARE",
		},
	}
}

func (c BatchAssetActionReqAction) Value() string {
	return c.value
}

func (c BatchAssetActionReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAssetActionReqAction) UnmarshalJSON(b []byte) error {
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
