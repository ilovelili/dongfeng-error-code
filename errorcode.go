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
	nutrition = "dongfeng.svc.nutrition.server"
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
	// GenericIPBlocked ip blocked
	GenericIPBlocked = new(generic, "E0006", "ip blocked", http.StatusForbidden)
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
	// CoreProxyBadFormatAttendanceFile attendance file is bad formatted
	CoreProxyBadFormatAttendanceFile = new(coreproxy, "E1008", "attendance file is bad formatted", http.StatusBadRequest)
	// CoreProxyInvalidAttendanceRequest invalid attendance request
	CoreProxyInvalidAttendanceRequest = new(coreproxy, "E1009", "invalid attendance request", http.StatusBadRequest)
	// CoreProxyFailedToReadPhysiqueFile failed to read physique file
	CoreProxyFailedToReadPhysiqueFile = new(coreproxy, "E1010", "failed to read physique file", http.StatusBadRequest)
	// CoreProxyBadFormatPhysiqueFile physique file is bad formatted
	CoreProxyBadFormatPhysiqueFile = new(coreproxy, "E1011", "physique file is bad formatted", http.StatusBadRequest)
	// CoreProxyBadFormatRecipeFile recipe file is bad formatted
	CoreProxyBadFormatRecipeFile = new(coreproxy, "E1012", "recipe file is bad formatted", http.StatusBadRequest)
	// CoreProxyInvalidUpdateRecipeRequestBody invalid recipe update request
	CoreProxyInvalidUpdateRecipeRequestBody = new(coreproxy, "E1013", "invalid recipe update request", http.StatusBadRequest)
	// CoreProxyInvalidUpdateIngredientRequestBody invalid ingredient update request
	CoreProxyInvalidUpdateIngredientRequestBody = new(coreproxy, "E1014", "invalid ingredient update request", http.StatusBadRequest)
	// CoreProxyInvalidGetMenuRequest invalid get menu request
	CoreProxyInvalidGetMenuRequest = new(coreproxy, "E1015", "invalid get menu request", http.StatusBadRequest)
	// CoreProxyInvalidGetProcurementRequest invalid get procurement request
	CoreProxyInvalidGetProcurementRequest = new(coreproxy, "E1016", "invalid get procurement request", http.StatusBadRequest)
	// CoreProxyInvalidUpdateNotificationRequestBody invalid update request
	CoreProxyInvalidUpdateNotificationRequestBody = new(coreproxy, "E1017", "invalid notification update request", http.StatusBadRequest)
	// CoreProxyInvalidClassListFile invalid class list file
	CoreProxyInvalidClassListFile = new(coreproxy, "E1018", "invalid class list file", http.StatusBadRequest)
	// CoreProxyInvalidTeacherListFile invalid teacher list file
	CoreProxyInvalidTeacherListFile = new(coreproxy, "E1019", "invalid teacher list file", http.StatusBadRequest)
	// CoreProxyInvalidClassNamelistFile invalid class namelist file
	CoreProxyInvalidClassNamelistFile = new(coreproxy, "E1020", "invalid class namelist file", http.StatusBadRequest)
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
	// CoreInvalidClass invalid class
	CoreInvalidClass = new(core, "E2005", "invalid class", http.StatusBadRequest)
	// CoreFailedToGetNamelist failed to get namelist
	CoreFailedToGetNamelist = new(core, "E2006", "failed to get namelist", http.StatusInternalServerError)
	// CoreFailedToUpdateNamelist failed to update namelist
	CoreFailedToUpdateNamelist = new(core, "E2007", "failed to update namelist", http.StatusInternalServerError)
	// CoreFailedToGetTeacherlist failed to get teacher list
	CoreFailedToGetTeacherlist = new(core, "E2008", "failed to get teacherlist", http.StatusInternalServerError)
	// CoreFailedToUpdateTeacherlist failed to update teacher list
	CoreFailedToUpdateTeacherlist = new(core, "E2009", "failed to update teacherlist", http.StatusInternalServerError)
	// CoreFailedToGetClasses failed to get classes
	CoreFailedToGetClasses = new(core, "E2010", "failed to get classes", http.StatusInternalServerError)
	// CoreFailedToUpdateClasses failed to update classes
	CoreFailedToUpdateClasses = new(core, "E2011", "failed to update classes", http.StatusInternalServerError)
)

// Nutrition
var (
	// NutritionNoUser user not found
	NutritionNoUser = new(nutrition, "E5000", "no user", http.StatusBadRequest)
	// NutritionFailedToGetIngredient failed to get ingredient
	NutritionFailedToGetIngredient = new(nutrition, "E5001", "failed to get ingredient", http.StatusInternalServerError)
	// NutritionFailedToSaveIngredient failed to save ingredient
	NutritionFailedToSaveIngredient = new(nutrition, "E5002", "failed to save ingredient", http.StatusInternalServerError)
	// NutritionFailedToSaveRecipe failed to save recipe
	NutritionFailedToSaveRecipe = new(nutrition, "E5003", "failed to save recipe", http.StatusInternalServerError)
	// NutritionFailedToGetMenu failed to get menu
	NutritionFailedToGetMenu = new(nutrition, "E5004", "failed to get menu", http.StatusInternalServerError)
	// NutritionFailedToGetRecipe failed to get recipe
	NutritionFailedToGetRecipe = new(nutrition, "E5005", "failed to get recipe", http.StatusInternalServerError)
	// NutritionFailedToGetProcurement failed to get procurements
	NutritionFailedToGetProcurement = new(nutrition, "E5006", "failed to get procurements", http.StatusInternalServerError)
)

// Error implements the error interface.
type Error struct {
	ID         string `json:"id"`
	Code       int64  `json:"code"`
	CustomCode string `json:"custom_code"`
	Detail     string `json:"detail"`
	Status     string `json:"status"`
}

// new private constructor
func new(id, customcode, detail string, code int64) *Error {
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

func newError(id, detail, customcode string, httpcode int64) error {
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
