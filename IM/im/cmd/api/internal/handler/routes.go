// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	push "backend/service/im/cmd/api/internal/handler/push"
	"backend/service/im/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/im/push",
				Handler: push.PushHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/im/pushroom",
				Handler: push.PushroomHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/im/count",
				Handler: push.CountHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/im/getroominfo",
				Handler: push.GetroominfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/im/v1"),
	)
}
