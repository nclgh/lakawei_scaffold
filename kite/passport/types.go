package passport

import (
	"encoding/json"
	"lakawei/lakawei_scaffold/kite/kite_common"
)

type CreateSessionRequest struct {
	UserId int64 `json:"user_id"`
}

func (r *CreateSessionRequest) Marshal() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func UnmarshalCreateSessionRequest(r string) (req *CreateSessionRequest, err error) {
	req = &CreateSessionRequest{}
	err = json.Unmarshal([]byte(r), req)
	return req, err
}

type CreateSessionResponse struct {
	Code kite_common.RspCode `json:"code"`
	Msg  string              `json:"msg"`
}

func GetCreateSessionResponse(code kite_common.RspCode) *CreateSessionResponse {
	rsp := &CreateSessionResponse{
		Code: code,
		Msg:  kite_common.UnknownCodeMsg,
	}
	msg, ok := kite_common.CodeMsg[code]
	if ok {
		rsp.Msg = msg
	}
	return rsp
}
