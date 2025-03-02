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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// AccountGetAuthorizations
// account.getAuthorizations#e320c158 = account.Authorizations;
func (c *PassportCore) AccountGetAuthorizations(in *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	_ = in

	authorizations, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionGetAuthorizations(c.ctx, &authsession.TLAuthsessionGetAuthorizations{
		UserId:           c.MD.UserId,
		ExcludeAuthKeyId: c.MD.PermAuthKeyId,
	})
	if err != nil {
		c.Logger.Errorf("account.getAuthorizations - error: %v", err)
		return nil, err
	}

	ttl, err := c.svcCtx.Dao.UserClient.UserGetAuthorizationTTL(c.ctx, &userpb.TLUserGetAuthorizationTTL{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("account.getAuthorizations - error: %v", err)
		return nil, err
	}

	authorizations.AuthorizationTtlDays = ttl.Days
	return authorizations, nil
}
