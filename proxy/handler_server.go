package proxy

import (
	"errors"

	"github.com/rueian/pgbroker/message"
)

type HandleAuthenticationOk func(ctx *Metadata, msg *message.AuthenticationOk) (*message.AuthenticationOk, error)
type HandleAuthenticationKerberosV5 func(ctx *Metadata, msg *message.AuthenticationKerberosV5) (*message.AuthenticationKerberosV5, error)
type HandleAuthenticationCleartextPassword func(ctx *Metadata, msg *message.AuthenticationCleartextPassword) (*message.AuthenticationCleartextPassword, error)
type HandleAuthenticationMD5Password func(ctx *Metadata, msg *message.AuthenticationMD5Password) (*message.AuthenticationMD5Password, error)
type HandleAuthenticationSCMCredential func(ctx *Metadata, msg *message.AuthenticationSCMCredential) (*message.AuthenticationSCMCredential, error)
type HandleAuthenticationGSS func(ctx *Metadata, msg *message.AuthenticationGSS) (*message.AuthenticationGSS, error)
type HandleAuthenticationSSPI func(ctx *Metadata, msg *message.AuthenticationSSPI) (*message.AuthenticationSSPI, error)
type HandleAuthenticationGSSContinue func(ctx *Metadata, msg *message.AuthenticationGSSContinue) (*message.AuthenticationGSSContinue, error)
type HandleAuthenticationSASL func(ctx *Metadata, msg *message.AuthenticationSASL) (*message.AuthenticationSASL, error)
type HandleAuthenticationSASLContinue func(ctx *Metadata, msg *message.AuthenticationSASLContinue) (*message.AuthenticationSASLContinue, error)
type HandleAuthenticationSASLFinal func(ctx *Metadata, msg *message.AuthenticationSASLFinal) (*message.AuthenticationSASLFinal, error)
type HandleBackendKeyData func(ctx *Metadata, msg *message.BackendKeyData) (*message.BackendKeyData, error)
type HandleBindComplete func(ctx *Metadata, msg *message.BindComplete) (*message.BindComplete, error)
type HandleCloseComplete func(ctx *Metadata, msg *message.CloseComplete) (*message.CloseComplete, error)
type HandleCommandComplete func(ctx *Metadata, msg *message.CommandComplete) (*message.CommandComplete, error)
type HandleCopyInResponse func(ctx *Metadata, msg *message.CopyInResponse) (*message.CopyInResponse, error)
type HandleCopyOutResponse func(ctx *Metadata, msg *message.CopyOutResponse) (*message.CopyOutResponse, error)
type HandleCopyBothResponse func(ctx *Metadata, msg *message.CopyBothResponse) (*message.CopyBothResponse, error)
type HandleDataRow func(ctx *Metadata, msg *message.DataRow) (*message.DataRow, error)
type HandleEmptyQueryResponse func(ctx *Metadata, msg *message.EmptyQueryResponse) (*message.EmptyQueryResponse, error)
type HandleErrorResponse func(ctx *Metadata, msg *message.ErrorResponse) (*message.ErrorResponse, error)
type HandleFunctionCallResponse func(ctx *Metadata, msg *message.FunctionCallResponse) (*message.FunctionCallResponse, error)
type HandleNegotiateProtocolVersion func(ctx *Metadata, msg *message.NegotiateProtocolVersion) (*message.NegotiateProtocolVersion, error)
type HandleNoData func(ctx *Metadata, msg *message.NoData) (*message.NoData, error)
type HandleNoticeResponse func(ctx *Metadata, msg *message.NoticeResponse) (*message.NoticeResponse, error)
type HandleNotificationResponse func(ctx *Metadata, msg *message.NotificationResponse) (*message.NotificationResponse, error)
type HandleParameterDescription func(ctx *Metadata, msg *message.ParameterDescription) (*message.ParameterDescription, error)
type HandleParameterStatus func(ctx *Metadata, msg *message.ParameterStatus) (*message.ParameterStatus, error)
type HandleParseComplete func(ctx *Metadata, msg *message.ParseComplete) (*message.ParseComplete, error)
type HandlePortalSuspended func(ctx *Metadata, msg *message.PortalSuspended) (*message.PortalSuspended, error)
type HandleReadyForQuery func(ctx *Metadata, msg *message.ReadyForQuery) (*message.ReadyForQuery, error)
type HandleRowDescription func(ctx *Metadata, msg *message.RowDescription) (*message.RowDescription, error)

type HandleAuthentication struct {
	HandleAuthenticationOk                HandleAuthenticationOk
	HandleAuthenticationKerberosV5        HandleAuthenticationKerberosV5
	HandleAuthenticationCleartextPassword HandleAuthenticationCleartextPassword
	HandleAuthenticationMD5Password       HandleAuthenticationMD5Password
	HandleAuthenticationSCMCredential     HandleAuthenticationSCMCredential
	HandleAuthenticationGSS               HandleAuthenticationGSS
	HandleAuthenticationSSPI              HandleAuthenticationSSPI
	HandleAuthenticationGSSContinue       HandleAuthenticationGSSContinue
	HandleAuthenticationSASL              HandleAuthenticationSASL
	HandleAuthenticationSASLContinue      HandleAuthenticationSASLContinue
	HandleAuthenticationSASLFinal         HandleAuthenticationSASLFinal
}

type ServerMessageHandlers map[byte]MessageHandler

func (s *ServerMessageHandlers) SetHandleAuthentication(h HandleAuthentication) {
	(*s)['R'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		switch m := message.ReadAuthentication(raw).(type) {
		case *message.AuthenticationOk:
			ctx.AuthPhase = PhaseOK
			if h.HandleAuthenticationOk == nil {
				return m, nil
			}
			return h.HandleAuthenticationOk(ctx, m)
		case *message.AuthenticationKerberosV5:
			if h.HandleAuthenticationKerberosV5 == nil {
				return m, nil
			}
			return h.HandleAuthenticationKerberosV5(ctx, m)
		case *message.AuthenticationCleartextPassword:
			if h.HandleAuthenticationCleartextPassword == nil {
				return m, nil
			}
			return h.HandleAuthenticationCleartextPassword(ctx, m)
		case *message.AuthenticationMD5Password:
			if h.HandleAuthenticationMD5Password == nil {
				return m, nil
			}
			return h.HandleAuthenticationMD5Password(ctx, m)
		case *message.AuthenticationSCMCredential:
			if h.HandleAuthenticationSCMCredential == nil {
				return m, nil
			}
			return h.HandleAuthenticationSCMCredential(ctx, m)
		case *message.AuthenticationGSS:
			ctx.AuthPhase = PhaseGSS
			if h.HandleAuthenticationGSS == nil {
				return m, nil
			}
			return h.HandleAuthenticationGSS(ctx, m)
		case *message.AuthenticationSSPI:
			if h.HandleAuthenticationSSPI == nil {
				return m, nil
			}
			return h.HandleAuthenticationSSPI(ctx, m)
		case *message.AuthenticationGSSContinue:
			if h.HandleAuthenticationGSSContinue == nil {
				return m, nil
			}
			return h.HandleAuthenticationGSSContinue(ctx, m)
		case *message.AuthenticationSASL:
			ctx.AuthPhase = PhaseSASLInit
			if h.HandleAuthenticationSASL == nil {
				return m, nil
			}
			return h.HandleAuthenticationSASL(ctx, m)
		case *message.AuthenticationSASLContinue:
			ctx.AuthPhase = PhaseSASL
			if h.HandleAuthenticationSASLContinue == nil {
				return m, nil
			}
			return h.HandleAuthenticationSASLContinue(ctx, m)
		case *message.AuthenticationSASLFinal:
			if h.HandleAuthenticationSASLFinal == nil {
				return m, nil
			}
			return h.HandleAuthenticationSASLFinal(ctx, m)
		}
		return nil, errors.New("fail to cast authentication message")
	}
}
func (s *ServerMessageHandlers) SetHandleBackendKeyData(h HandleBackendKeyData) {
	if h == nil {
		return
	}
	(*s)['K'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadBackendKeyData(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleBindComplete(h HandleBindComplete) {
	if h == nil {
		return
	}
	(*s)['2'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadBindComplete(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleCloseComplete(h HandleCloseComplete) {
	if h == nil {
		return
	}
	(*s)['3'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadCloseComplete(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleCommandComplete(h HandleCommandComplete) {
	if h == nil {
		return
	}
	(*s)['C'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadCommandComplete(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleCopyInResponse(h HandleCopyInResponse) {
	if h == nil {
		return
	}
	(*s)['G'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadCopyInResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleCopyOutResponse(h HandleCopyOutResponse) {
	if h == nil {
		return
	}
	(*s)['H'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadCopyOutResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleCopyBothResponse(h HandleCopyBothResponse) {
	if h == nil {
		return
	}
	(*s)['W'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadCopyBothResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleDataRow(h HandleDataRow) {
	if h == nil {
		return
	}
	(*s)['D'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadDataRow(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleEmptyQueryResponse(h HandleEmptyQueryResponse) {
	if h == nil {
		return
	}
	(*s)['I'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadEmptyQueryResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleErrorResponse(h HandleErrorResponse) {
	if h == nil {
		return
	}
	(*s)['E'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadErrorResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleFunctionCallResponse(h HandleFunctionCallResponse) {
	if h == nil {
		return
	}
	(*s)['V'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadFunctionCallResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleNegotiateProtocolVersion(h HandleNegotiateProtocolVersion) {
	if h == nil {
		return
	}
	(*s)['v'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadNegotiateProtocolVersion(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleNoData(h HandleNoData) {
	if h == nil {
		return
	}
	(*s)['n'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadNoData(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleNoticeResponse(h HandleNoticeResponse) {
	if h == nil {
		return
	}
	(*s)['N'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadNoticeResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleNotificationResponse(h HandleNotificationResponse) {
	if h == nil {
		return
	}
	(*s)['A'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadNotificationResponse(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleParameterDescription(h HandleParameterDescription) {
	if h == nil {
		return
	}
	(*s)['t'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadParameterDescription(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleParameterStatus(h HandleParameterStatus) {
	if h == nil {
		return
	}
	(*s)['S'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadParameterStatus(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleParseComplete(h HandleParseComplete) {
	if h == nil {
		return
	}
	(*s)['1'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadParseComplete(raw))
	}
}
func (s *ServerMessageHandlers) SetHandlePortalSuspended(h HandlePortalSuspended) {
	if h == nil {
		return
	}
	(*s)['s'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadPortalSuspended(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleReadyForQuery(h HandleReadyForQuery) {
	if h == nil {
		return
	}
	(*s)['Z'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadReadyForQuery(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleRowDescription(h HandleRowDescription) {
	if h == nil {
		return
	}
	(*s)['T'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadRowDescription(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleCopyData(h HandleCopyData) {
	if h == nil {
		return
	}
	(*s)['d'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadCopyData(raw))
	}
}
func (s *ServerMessageHandlers) SetHandleCopyDone(h HandleCopyDone) {
	if h == nil {
		return
	}
	(*s)['c'] = func(ctx *Metadata, raw []byte) (message.Reader, error) {
		return h(ctx, message.ReadCopyDone(raw))
	}
}
