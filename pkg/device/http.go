// Copyright 2023 Tiptopsoft, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package device

import (
	"github.com/gin-gonic/gin"
	"github.com/tiptopsoft/fvpn/pkg/util"
)

var (
	PREFIX = "/api/v1/"
)

func (n *Node) HttpServer() error {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.POST(PREFIX+"join", n.joinNet())
	server.POST(PREFIX+"leave", n.leaveNet())

	return server.Run(n.cfg.HttpListenStr())
}

func (n *Node) joinNet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req JoinRequest
		err := ctx.Bind(&req)

		if err != nil {
			ctx.JSON(500, HttpError(err.Error()))
			return
		}

		if req.CIDR != "" {
			err = n.netCtl.JoinNet(util.UCTL.UserId, req.CIDR)
			if err != nil {
				ctx.JSON(500, HttpError(err.Error()))
				return
			}
		} else {
			ctx.JSON(500, HttpError("cidr is nil"))
			return
		}

		resp := &JoinResponse{
			IP:   n.device.IPToString(),
			Name: n.device.Name(),
		}

		//userId替换
		ctx.JSON(200, HttpOK(resp))
	}
}

func (n *Node) leaveNet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req JoinRequest
		err := ctx.Bind(&req)

		if err != nil {
			ctx.JSON(500, HttpError(err.Error()))
			return
		}

		if req.CIDR != "" {
			err = n.netCtl.LeaveNet(util.UCTL.UserId, req.CIDR)
			if err != nil {
				ctx.JSON(500, HttpError(err.Error()))
				return
			}
		} else {
			ctx.JSON(500, HttpError("cidr is nil"))
			return
		}

		resp := new(LeaveResponse)
		resp.IP = n.device.IPToString()
		resp.Name = n.device.Name()

		ctx.JSON(200, HttpOK(resp))
	}
}
