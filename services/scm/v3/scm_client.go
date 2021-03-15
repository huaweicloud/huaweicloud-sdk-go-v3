package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/scm/v3/model"
)

type ScmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewScmClient(hcClient *http_client.HcHttpClient) *ScmClient {
	return &ScmClient{HcClient: hcClient}
}

func ScmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//删除证书实例，即将证书资源从华为云系统中删除。
func (c *ScmClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

//导出证书。
func (c *ScmClient) ExportCertificate(request *model.ExportCertificateRequest) (*model.ExportCertificateResponse, error) {
	requestDef := GenReqDefForExportCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportCertificateResponse), nil
	}
}

//导入证书到SCM服务管理。
func (c *ScmClient) ImportCertificate(request *model.ImportCertificateRequest) (*model.ImportCertificateResponse, error) {
	requestDef := GenReqDefForImportCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportCertificateResponse), nil
	}
}

//根据证书名称或绑定域名查询证书列表。
func (c *ScmClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

//推送SSL证书到弹性负载均衡（Elastic Load Balance，简称ELB）、Web应用防火墙（Web Application Firewall，WAF）、CDN（Content Delivery Network，内容分发网络）等其它华为云产品中。
func (c *ScmClient) PushCertificate(request *model.PushCertificateRequest) (*model.PushCertificateResponse, error) {
	requestDef := GenReqDefForPushCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PushCertificateResponse), nil
	}
}

//查询某张证书的详细信息。
func (c *ScmClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}
