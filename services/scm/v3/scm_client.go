package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/scm/v3/model"
)

type ScmClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewScmClient(hcClient *httpclient.HcHttpClient) *ScmClient {
	return &ScmClient{HcClient: hcClient}
}

func ScmClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials,basic.Credentials")
	return builder
}

// ApplyCertificate 申请证书
//
// 申请证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) ApplyCertificate(request *model.ApplyCertificateRequest) (*model.ApplyCertificateResponse, error) {
	requestDef := GenReqDefForApplyCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyCertificateResponse), nil
	}
}

// ApplyCertificateInvoker 申请证书
func (c *ScmClient) ApplyCertificateInvoker(request *model.ApplyCertificateRequest) *ApplyCertificateInvoker {
	requestDef := GenReqDefForApplyCertificate()
	return &ApplyCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchPushCertificate 批量推送证书
//
// 批量推送SSL证书到弹性负载均衡（Elastic Load Balance，简称ELB）、Web应用防火墙（Web Application Firewall，WAF）、CDN（Content Delivery Network，内容分发网络）等其它华为云产品中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) BatchPushCertificate(request *model.BatchPushCertificateRequest) (*model.BatchPushCertificateResponse, error) {
	requestDef := GenReqDefForBatchPushCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchPushCertificateResponse), nil
	}
}

// BatchPushCertificateInvoker 批量推送证书
func (c *ScmClient) BatchPushCertificateInvoker(request *model.BatchPushCertificateRequest) *BatchPushCertificateInvoker {
	requestDef := GenReqDefForBatchPushCertificate()
	return &BatchPushCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
//
// 删除证书实例，即将证书资源从系统中删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// DeleteCertificateInvoker 删除证书
func (c *ScmClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeployCertificate 部署证书
//
// 部署SSL证书到弹性负载均衡（Elastic Load Balance，简称ELB）、Web应用防火墙（Web Application Firewall，WAF）、CDN（Content Delivery Network，内容分发网络）等其它华为云产品中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) DeployCertificate(request *model.DeployCertificateRequest) (*model.DeployCertificateResponse, error) {
	requestDef := GenReqDefForDeployCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeployCertificateResponse), nil
	}
}

// DeployCertificateInvoker 部署证书
func (c *ScmClient) DeployCertificateInvoker(request *model.DeployCertificateRequest) *DeployCertificateInvoker {
	requestDef := GenReqDefForDeployCertificate()
	return &DeployCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportCertificate 导出证书
//
// 导出证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) ExportCertificate(request *model.ExportCertificateRequest) (*model.ExportCertificateResponse, error) {
	requestDef := GenReqDefForExportCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportCertificateResponse), nil
	}
}

// ExportCertificateInvoker 导出证书
func (c *ScmClient) ExportCertificateInvoker(request *model.ExportCertificateRequest) *ExportCertificateInvoker {
	requestDef := GenReqDefForExportCertificate()
	return &ExportCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportCertificate 导入证书
//
// 导入证书到SCM服务管理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) ImportCertificate(request *model.ImportCertificateRequest) (*model.ImportCertificateResponse, error) {
	requestDef := GenReqDefForImportCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportCertificateResponse), nil
	}
}

// ImportCertificateInvoker 导入证书
func (c *ScmClient) ImportCertificateInvoker(request *model.ImportCertificateRequest) *ImportCertificateInvoker {
	requestDef := GenReqDefForImportCertificate()
	return &ImportCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 查询证书列表
//
// 根据证书名称或绑定域名查询证书列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

// ListCertificatesInvoker 查询证书列表
func (c *ScmClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeployedResources 查询已部署资源
//
// 查询证书已部署的具体资源。针对已签发和上传的非国密证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) ListDeployedResources(request *model.ListDeployedResourcesRequest) (*model.ListDeployedResourcesResponse, error) {
	requestDef := GenReqDefForListDeployedResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeployedResourcesResponse), nil
	}
}

// ListDeployedResourcesInvoker 查询已部署资源
func (c *ScmClient) ListDeployedResourcesInvoker(request *model.ListDeployedResourcesRequest) *ListDeployedResourcesInvoker {
	requestDef := GenReqDefForListDeployedResources()
	return &ListDeployedResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PushCertificate 推送证书
//
// 推送SSL证书到弹性负载均衡（Elastic Load Balance，简称ELB）、Web应用防火墙（Web Application Firewall，WAF）、CDN（Content Delivery Network，内容分发网络）等其它华为云产品中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) PushCertificate(request *model.PushCertificateRequest) (*model.PushCertificateResponse, error) {
	requestDef := GenReqDefForPushCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushCertificateResponse), nil
	}
}

// PushCertificateInvoker 推送证书
func (c *ScmClient) PushCertificateInvoker(request *model.PushCertificateRequest) *PushCertificateInvoker {
	requestDef := GenReqDefForPushCertificate()
	return &PushCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 获取证书详情
//
// 查询某张证书的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

// ShowCertificateInvoker 获取证书详情
func (c *ScmClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeCertificate 购买SSL证书
//
// 购买SSL证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) SubscribeCertificate(request *model.SubscribeCertificateRequest) (*model.SubscribeCertificateResponse, error) {
	requestDef := GenReqDefForSubscribeCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeCertificateResponse), nil
	}
}

// SubscribeCertificateInvoker 购买SSL证书
func (c *ScmClient) SubscribeCertificateInvoker(request *model.SubscribeCertificateRequest) *SubscribeCertificateInvoker {
	requestDef := GenReqDefForSubscribeCertificate()
	return &SubscribeCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnsubscribeCertificate 退订证书
//
// 退订证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ScmClient) UnsubscribeCertificate(request *model.UnsubscribeCertificateRequest) (*model.UnsubscribeCertificateResponse, error) {
	requestDef := GenReqDefForUnsubscribeCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnsubscribeCertificateResponse), nil
	}
}

// UnsubscribeCertificateInvoker 退订证书
func (c *ScmClient) UnsubscribeCertificateInvoker(request *model.UnsubscribeCertificateRequest) *UnsubscribeCertificateInvoker {
	requestDef := GenReqDefForUnsubscribeCertificate()
	return &UnsubscribeCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
