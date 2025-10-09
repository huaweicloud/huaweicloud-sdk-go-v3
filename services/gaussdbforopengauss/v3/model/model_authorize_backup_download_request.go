package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthorizeBackupDownloadRequest Request Object
type AuthorizeBackupDownloadRequest struct {

	// **参数解释**: 用户Token。 通过调用IAM服务[获取用户token](https://support.huaweicloud.com/intl/zh-cn/api-iam/iam_30_0001.html)。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	XLanguage *AuthorizeBackupDownloadRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 备份ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	BackupId string `json:"backup_id"`
}

func (o AuthorizeBackupDownloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeBackupDownloadRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeBackupDownloadRequest", string(data)}, " ")
}

type AuthorizeBackupDownloadRequestXLanguage struct {
	value string
}

type AuthorizeBackupDownloadRequestXLanguageEnum struct {
	ZH_CN AuthorizeBackupDownloadRequestXLanguage
	EN_US AuthorizeBackupDownloadRequestXLanguage
}

func GetAuthorizeBackupDownloadRequestXLanguageEnum() AuthorizeBackupDownloadRequestXLanguageEnum {
	return AuthorizeBackupDownloadRequestXLanguageEnum{
		ZH_CN: AuthorizeBackupDownloadRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: AuthorizeBackupDownloadRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c AuthorizeBackupDownloadRequestXLanguage) Value() string {
	return c.value
}

func (c AuthorizeBackupDownloadRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeBackupDownloadRequestXLanguage) UnmarshalJSON(b []byte) error {
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
