// Package errorcode defines the customzied error code.
package errorcode

import (
	"encoding/json"
	"net/http"
)

// ids
const (
	pipe      = "passthru"
	generic   = "dongfeng.svc.generic"
	core      = "dongfeng.svc.core.server"
	coreproxy = "dongfeng.svc.core.proxy"
)

var (
	// Pipe means to passthru the customcode from other microservices
	Pipe = new(pipe, "PASSTHRU", "passthru", http.StatusBadGateway)
)

// Generic errors
var (
	// GenericHeIsDeadJim healthcheck dead
	GenericHeIsDeadJim = new(generic, "E0000", "he is dead, Jim", http.StatusBadGateway)
	// GenericNotAuthorized invalid token
	GenericNotAuthorized = new(generic, "E0001", "not authorized", http.StatusUnauthorized)
	// GenericInvalidToken invalid token
	GenericInvalidToken = new(generic, "E0002", "invalid token", http.StatusBadRequest)
	// GenericTokenExpired token expired
	GenericTokenExpired = new(generic, "E0003", "token expired", http.StatusBadRequest)
	// GenericInsufficientPrivileges insufficient privileges
	GenericInsufficientPrivileges = new(generic, "E0004", "insufficient privileges", http.StatusForbidden)
	// GenericInvalidMetaData invalid meta data
	GenericInvalidMetaData = new(generic, "E0005", "invalid meta data", http.StatusBadRequest)
)

// Core-Proxy
var (
	// CoreProxyFailedToGetSession filed to get session from redis
	CoreProxyFailedToGetSession = new(coreproxy, "E1001", "failed to get session", http.StatusInternalServerError)
	// CoreProxyFailedToGetSession failed to save session to redis
	CoreProxyFailedToSaveSession = new(coreproxy, "E1002", "failed to save session", http.StatusInternalServerError)
	// CoreProxyFailedToReadAvatarFile failed to read avatar file
	CoreProxyFailedToReadAvatarFile = new(coreproxy, "E1003", "failed to read avatar file", http.StatusBadRequest)
	// CoreProxyFailedToUploadAvatar failed to upload avatar file
	CoreProxyFailedToUploadAvatar = new(coreproxy, "E1004", "failed to upload avatar file", http.StatusInternalServerError)
	// CoreProxyUnsupportedMimeType unsupported mime type
	CoreProxyUnsupportedMimeType = new(coreproxy, "E1005", "unsupported mime type", http.StatusBadRequest)
	// CoreProxyInvalidUpdateUserRequestBody invalid update request
	CoreProxyInvalidUpdateUserRequestBody = new(coreproxy, "E1006", "invalid user update request", http.StatusBadRequest)
	// CoreProxyFailedToReadAttendanceFile failed to read attendance file
	CoreProxyFailedToReadAttendanceFile = new(coreproxy, "E1007", "failed to read attendance file", http.StatusBadRequest)
)

// Core
var (
	// CoreFailedToSaveUser failed to save user
	CoreFailedToSaveUser = new(core, "E2000", "failed to save user", http.StatusInternalServerError)
	// CoreFailedToGetNotification failed to get notificaiton
	CoreFailedToGetNotification = new(core, "E2001", "failed to get notificaiton", http.StatusInternalServerError)
	// CoreFailedToGetFriends failed to get friends
	CoreFailedToGetFriends = new(core, "E2002", "failed to get friends", http.StatusInternalServerError)
	// CoreNoUser user not found
	CoreNoUser = new(core, "E2003", "no user", http.StatusBadRequest)
	// CoreInvalidUpdateUserRequest invalid update user request
	CoreInvalidUpdateUserRequest = new(core, "E2004", "invalid update user request", http.StatusBadRequest)
)

// Error implements the error interface.
type Error struct {
	ID         string `json:"id"`
	Code       int32  `json:"code"`
	CustomCode string `json:"custom_code"`
	Detail     string `json:"detail"`
	Status     string `json:"status"`
}

// new private constructor
func new(id, customcode, detail string, code int32) *Error {
	return &Error{
		ID:         id,
		CustomCode: customcode,
		Detail:     detail,
		Code:       code,
	}
}

// NewError new rpc error
func (e *Error) NewError(detail ...string) error {
	if len(detail) == 1 {
		e.Detail = detail[0]
	}
	return newError(e.ID, e.Detail, e.CustomCode, e.Code)
}

// Error error to string
func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func newError(id, detail, customcode string, httpcode int32) error {
	err := parse(detail)
	code := err.Code
	ccode := err.CustomCode

	// is errorcode is pipe, passthru
	if customcode != Pipe.CustomCode {
		code = httpcode
		ccode = customcode
	}

	return &Error{
		ID:         id,
		Code:       code,
		CustomCode: ccode,
		Detail:     detail,
		Status:     http.StatusText(int(code)),
	}
}

// parse tries to parse a JSON string into an error. If that fails, it will set the given string as the error detail.
func parse(err string) *Error {
	e := &Error{}
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Detail = err
	}
	return e
}
