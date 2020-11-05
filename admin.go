package admin

import (
	`fmt`
	`net/http`

	`github.com/class100/yunke-core`
	`github.com/dgrijalva/jwt-go`
	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

type (
	// ClientAdmin 云视课堂管理客户端
	ClientAdmin struct {
		// Id 产品编号
		Id int64
		// Url 地址
		Url string
		// Secret 通信秘钥
		Secret string
		// SigningMethod 加密方法
		SigningMethod string
		// AuthScheme 授权码，在Header里面
		AuthScheme string
	}
)

func (ca *ClientAdmin) request(
	path core.ApiPath,
	method gox.HttpMethod,
	params interface{}, pathParams map[string]string,
	rsp interface{},
) (err error) {
	var (
		adminRsp           *resty.Response
		authToken          string
		expectedStatusCode int
	)

	if authToken, err = token(DefaultAdminDomain, jwt.GetSigningMethod(ca.SigningMethod), ca.Secret); nil != err {
		return
	}

	req := NewResty().SetResult(rsp).SetHeader(gox.HeaderAuthorization, fmt.Sprintf("%s %s", ca.AuthScheme, authToken))
	// 注入路径参数
	if 0 != len(pathParams) {
		req = req.SetPathParams(pathParams)
	}

	// 修正请求地址为全路径
	url := fmt.Sprintf("%s/api/%s", ca.Url, path)

	switch method {
	case gox.HttpMethodGet:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetQueryParams(params.(map[string]string))
		}
		adminRsp, err = req.Get(url)
	case gox.HttpMethodPost:
		expectedStatusCode = http.StatusCreated

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Post(url)
	case gox.HttpMethodPut:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Put(url)
	case gox.HttpMethodDelete:
		expectedStatusCode = http.StatusNoContent

		if nil != params {
			req = req.SetBody(params)
		}
		adminRsp, err = req.Delete(url)
	}
	if nil != err {
		return
	}

	if nil == adminRsp {
		err = gox.NewCodeError(gox.ErrorCode(adminRsp.StatusCode()), "无返回数据", RestyStringBody(adminRsp))

		return
	}

	// 检查状态码
	if expectedStatusCode != adminRsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(adminRsp.StatusCode()), "请求服务器不符合预期", RestyStringBody(adminRsp))
	}

	return
}
