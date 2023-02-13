package result

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"zcloud/pkg/e"
	"zcloud/pkg/e/pb"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := e.ServerCommonError
		errmsg := e.MapErrMsg(e.ServerCommonError)
		causeErr := errors.Cause(err)             // err类型
		if e, ok := causeErr.(*e.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				for _, v := range gstatus.Details() {
					detail := v.(*pb.Error)
					grpcCode := uint32(detail.Code)
					if e.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
						errcode = grpcCode
						errmsg = detail.Msg
					}
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("[API-ERR]: %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
	}
}

func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := e.ServerCommonError
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)             // err类型
		if e, ok := causeErr.(*e.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if e.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("[GATEWAY-ERR]: %+v ", err)

		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", e.MapErrMsg(e.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(e.RequestParamError, errMsg))
}
