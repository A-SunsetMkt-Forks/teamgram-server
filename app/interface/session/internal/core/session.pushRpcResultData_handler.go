// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/session/session"
)

// SessionPushRpcResultData
// session.pushRpcResultData auth_key_id:long session_id:long client_req_msg_id:long rpc_result_data:bytes = Bool;
func (c *SessionCore) SessionPushRpcResultData(in *session.TLSessionPushRpcResultData) (*mtproto.Bool, error) {
	mainAuth := c.svcCtx.MainAuthMgr.GetMainAuthWrapper(in.PermAuthKeyId)
	if mainAuth == nil {
		err := fmt.Errorf("not found authKeyId(%s)", in)
		c.Logger.Errorf("session.pushRpcResultData - %v", err)
		return nil, err
	}
	_ = mainAuth.SyncRpcResultDataArrived(c.ctx, in.AuthKeyId, in.SessionId, in.ClientReqMsgId, in.RpcResultData)

	return mtproto.BoolTrue, nil
}
