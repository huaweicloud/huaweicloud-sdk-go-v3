package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/msgsms/v2/model"
)

type MsgsmsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewMsgsmsClient(hcClient *httpclient.HcHttpClient) *MsgsmsClient {
	return &MsgsmsClient{HcClient: hcClient}
}

func MsgsmsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateApp 创建短信应用
//
// 该接口用于用户创建应用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建短信应用
func (c *MsgsmsClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppDetails 查询短信应用
//
// 该接口用于用户查询已创建的应用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ListAppDetails(request *model.ListAppDetailsRequest) (*model.ListAppDetailsResponse, error) {
	requestDef := GenReqDefForListAppDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppDetailsResponse), nil
	}
}

// ListAppDetailsInvoker 查询短信应用
func (c *MsgsmsClient) ListAppDetailsInvoker(request *model.ListAppDetailsRequest) *ListAppDetailsInvoker {
	requestDef := GenReqDefForListAppDetails()
	return &ListAppDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApp 获取应用详情
//
// 该接口用于用户查询应用详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// ShowAppInvoker 获取应用详情
func (c *MsgsmsClient) ShowAppInvoker(request *model.ShowAppRequest) *ShowAppInvoker {
	requestDef := GenReqDefForShowApp()
	return &ShowAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppCount 查询应用数量
//
// 该接口用于用户查询应用使用的数量信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ShowAppCount(request *model.ShowAppCountRequest) (*model.ShowAppCountResponse, error) {
	requestDef := GenReqDefForShowAppCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppCountResponse), nil
	}
}

// ShowAppCountInvoker 查询应用数量
func (c *MsgsmsClient) ShowAppCountInvoker(request *model.ShowAppCountRequest) *ShowAppCountInvoker {
	requestDef := GenReqDefForShowAppCount()
	return &ShowAppCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApp 修改短信应用
//
// 该接口用于用户修改应用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

// UpdateAppInvoker 修改短信应用
func (c *MsgsmsClient) UpdateAppInvoker(request *model.UpdateAppRequest) *UpdateAppInvoker {
	requestDef := GenReqDefForUpdateApp()
	return &UpdateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSignature 创建短信签名
//
// 该接口用于用户创建签名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) CreateSignature(request *model.CreateSignatureRequest) (*model.CreateSignatureResponse, error) {
	requestDef := GenReqDefForCreateSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSignatureResponse), nil
	}
}

// CreateSignatureInvoker 创建短信签名
func (c *MsgsmsClient) CreateSignatureInvoker(request *model.CreateSignatureRequest) *CreateSignatureInvoker {
	requestDef := GenReqDefForCreateSignature()
	return &CreateSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSignature 删除短信签名
//
// 该接口用于用户删除已创建的签名信息息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) DeleteSignature(request *model.DeleteSignatureRequest) (*model.DeleteSignatureResponse, error) {
	requestDef := GenReqDefForDeleteSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSignatureResponse), nil
	}
}

// DeleteSignatureInvoker 删除短信签名
func (c *MsgsmsClient) DeleteSignatureInvoker(request *model.DeleteSignatureRequest) *DeleteSignatureInvoker {
	requestDef := GenReqDefForDeleteSignature()
	return &DeleteSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableSignature 申请激活签名
//
// 该接口用于用户申请激活签名信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) EnableSignature(request *model.EnableSignatureRequest) (*model.EnableSignatureResponse, error) {
	requestDef := GenReqDefForEnableSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableSignatureResponse), nil
	}
}

// EnableSignatureInvoker 申请激活签名
func (c *MsgsmsClient) EnableSignatureInvoker(request *model.EnableSignatureRequest) *EnableSignatureInvoker {
	requestDef := GenReqDefForEnableSignature()
	return &EnableSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSignatureDetails 查询签名信息
//
// 该接口用于用户查询已创建的短信签名信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ListSignatureDetails(request *model.ListSignatureDetailsRequest) (*model.ListSignatureDetailsResponse, error) {
	requestDef := GenReqDefForListSignatureDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSignatureDetailsResponse), nil
	}
}

// ListSignatureDetailsInvoker 查询签名信息
func (c *MsgsmsClient) ListSignatureDetailsInvoker(request *model.ListSignatureDetailsRequest) *ListSignatureDetailsInvoker {
	requestDef := GenReqDefForListSignatureDetails()
	return &ListSignatureDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSignature 获取签名详情
//
// 该接口用于用户查询签名详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ShowSignature(request *model.ShowSignatureRequest) (*model.ShowSignatureResponse, error) {
	requestDef := GenReqDefForShowSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSignatureResponse), nil
	}
}

// ShowSignatureInvoker 获取签名详情
func (c *MsgsmsClient) ShowSignatureInvoker(request *model.ShowSignatureRequest) *ShowSignatureInvoker {
	requestDef := GenReqDefForShowSignature()
	return &ShowSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSignatureFile 查询申请文件
//
// 该接口用于用户查询上传的文件信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ShowSignatureFile(request *model.ShowSignatureFileRequest) (*model.ShowSignatureFileResponse, error) {
	requestDef := GenReqDefForShowSignatureFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSignatureFileResponse), nil
	}
}

// ShowSignatureFileInvoker 查询申请文件
func (c *MsgsmsClient) ShowSignatureFileInvoker(request *model.ShowSignatureFileRequest) *ShowSignatureFileInvoker {
	requestDef := GenReqDefForShowSignatureFile()
	return &ShowSignatureFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSignature 修改短信签名
//
// 该接口用于用户更新签名信息，目前仅支持审核不通过的短信签名重新修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) UpdateSignature(request *model.UpdateSignatureRequest) (*model.UpdateSignatureResponse, error) {
	requestDef := GenReqDefForUpdateSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSignatureResponse), nil
	}
}

// UpdateSignatureInvoker 修改短信签名
func (c *MsgsmsClient) UpdateSignatureInvoker(request *model.UpdateSignatureRequest) *UpdateSignatureInvoker {
	requestDef := GenReqDefForUpdateSignature()
	return &UpdateSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadSignatureFile 上传申请文件
//
// 该接口用于用户上传文件信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) UploadSignatureFile(request *model.UploadSignatureFileRequest) (*model.UploadSignatureFileResponse, error) {
	requestDef := GenReqDefForUploadSignatureFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadSignatureFileResponse), nil
	}
}

// UploadSignatureFileInvoker 上传申请文件
func (c *MsgsmsClient) UploadSignatureFileInvoker(request *model.UploadSignatureFileRequest) *UploadSignatureFileInvoker {
	requestDef := GenReqDefForUploadSignatureFile()
	return &UploadSignatureFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplate 创建短信模板
//
// 该接口用于用户创建模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// CreateTemplateInvoker 创建短信模板
func (c *MsgsmsClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplate 删除短信模板
//
// 该接口用于用户删除已创建的模板信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// DeleteTemplateInvoker 删除短信模板
func (c *MsgsmsClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSendCountryDetails 查询发送国家
//
// 该接口用于用户查询短信发送的国家信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ListSendCountryDetails(request *model.ListSendCountryDetailsRequest) (*model.ListSendCountryDetailsResponse, error) {
	requestDef := GenReqDefForListSendCountryDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSendCountryDetailsResponse), nil
	}
}

// ListSendCountryDetailsInvoker 查询发送国家
func (c *MsgsmsClient) ListSendCountryDetailsInvoker(request *model.ListSendCountryDetailsRequest) *ListSendCountryDetailsInvoker {
	requestDef := GenReqDefForListSendCountryDetails()
	return &ListSendCountryDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplateDetails 查询短信模板
//
// 该接口用于用户查询已创建的模板信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ListTemplateDetails(request *model.ListTemplateDetailsRequest) (*model.ListTemplateDetailsResponse, error) {
	requestDef := GenReqDefForListTemplateDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateDetailsResponse), nil
	}
}

// ListTemplateDetailsInvoker 查询短信模板
func (c *MsgsmsClient) ListTemplateDetailsInvoker(request *model.ListTemplateDetailsRequest) *ListTemplateDetailsInvoker {
	requestDef := GenReqDefForListTemplateDetails()
	return &ListTemplateDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplateVarilableDetails 查询模板变量
//
// 该接口用于用户查询模板参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ListTemplateVarilableDetails(request *model.ListTemplateVarilableDetailsRequest) (*model.ListTemplateVarilableDetailsResponse, error) {
	requestDef := GenReqDefForListTemplateVarilableDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateVarilableDetailsResponse), nil
	}
}

// ListTemplateVarilableDetailsInvoker 查询模板变量
func (c *MsgsmsClient) ListTemplateVarilableDetailsInvoker(request *model.ListTemplateVarilableDetailsRequest) *ListTemplateVarilableDetailsInvoker {
	requestDef := GenReqDefForListTemplateVarilableDetails()
	return &ListTemplateVarilableDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplate 获取模板详情
//
// 该接口用于用户查询已创建的模板详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) ShowTemplate(request *model.ShowTemplateRequest) (*model.ShowTemplateResponse, error) {
	requestDef := GenReqDefForShowTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateResponse), nil
	}
}

// ShowTemplateInvoker 获取模板详情
func (c *MsgsmsClient) ShowTemplateInvoker(request *model.ShowTemplateRequest) *ShowTemplateInvoker {
	requestDef := GenReqDefForShowTemplate()
	return &ShowTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTemplate 修改短信模板
//
// 该接口用于用户修改模板信息，目前仅支持审核不通过的短信模板重新修改
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MsgsmsClient) UpdateTemplate(request *model.UpdateTemplateRequest) (*model.UpdateTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateResponse), nil
	}
}

// UpdateTemplateInvoker 修改短信模板
func (c *MsgsmsClient) UpdateTemplateInvoker(request *model.UpdateTemplateRequest) *UpdateTemplateInvoker {
	requestDef := GenReqDefForUpdateTemplate()
	return &UpdateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
