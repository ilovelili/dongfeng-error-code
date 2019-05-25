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
	// CoreProxyInvalidClassUploadFile invalid class list file
	CoreProxyInvalidClassUploadFile = new(coreproxy, "E1018", "invalid class upload file", http.StatusBadRequest)
	// CoreProxyInvalidTeacherUploadFile invalid teacher list file
	CoreProxyInvalidTeacherUploadFile = new(coreproxy, "E1019", "invalid teacher upload file", http.StatusBadRequest)
	// CoreProxyInvalidPupilUploadFile invalid pupil file
	CoreProxyInvalidPupilUploadFile = new(coreproxy, "E1020", "invalid pupil upload file", http.StatusBadRequest)
	// CoreProxyInvalidAttendanceUploadFile invalid attendance file
	CoreProxyInvalidAttendanceUploadFile = new(coreproxy, "E1021", "invalid attendance upload file", http.StatusBadRequest)
	// CoreProxyInvalidPhysiqueUploadFile invalid physique file
	CoreProxyInvalidPhysiqueUploadFile = new(coreproxy, "E1022", "invalid physique upload file", http.StatusBadRequest)
	// CoreProxyInvalidAttendanceUpdateRequestBody invalid attendance update request
	CoreProxyInvalidAttendanceUpdateRequestBody = new(coreproxy, "E1023", "invalid attendance update request", http.StatusBadRequest)
	// CoreProxyInvalidPupilUpdateRequestBody invalid pupil update request
	CoreProxyInvalidPupilUpdateRequestBody = new(coreproxy, "E1024", "invalid pupil update request", http.StatusBadRequest)
	// CoreProxyInvalidTeacherUpdateRequestBody invalid teacher update request
	CoreProxyInvalidTeacherUpdateRequestBody = new(coreproxy, "E1025", "invalid teacher update request", http.StatusBadRequest)
	// CoreProxyInvalidPhysiqueUpdateRequestBody invalid physique update request
	CoreProxyInvalidPhysiqueUpdateRequestBody = new(coreproxy, "E1026", "invalid physique update request", http.StatusBadRequest)
	// CoreProxyInvalidPhysiqueMasterRequest invalid physique master request
	CoreProxyInvalidPhysiqueMasterRequest = new(coreproxy, "E1027", "invalid physique master request", http.StatusBadRequest)
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
	// CoreFailedToGetPupils failed to get pupils
	CoreFailedToGetPupils = new(core, "E2006", "failed to get pupils", http.StatusInternalServerError)
	// CoreFailedToUpdatePupils failed to update pupils
	CoreFailedToUpdatePupils = new(core, "E2007", "failed to update pupils", http.StatusInternalServerError)
	// CoreFailedToGetTeachers failed to get teacher list
	CoreFailedToGetTeachers = new(core, "E2008", "failed to get teachers", http.StatusInternalServerError)
	// CoreFailedToUpdateTeachers failed to update teacher list
	CoreFailedToUpdateTeachers = new(core, "E2009", "failed to update teachers", http.StatusInternalServerError)
	// CoreFailedToGetClasses failed to get classes
	CoreFailedToGetClasses = new(core, "E2010", "failed to get classes", http.StatusInternalServerError)
	// CoreFailedToUpdateClasses failed to update classes
	CoreFailedToUpdateClasses = new(core, "E2011", "failed to update classes", http.StatusInternalServerError)
	// CoreFailedToGetAttendances failed to get attendances
	CoreFailedToGetAttendances = new(core, "E2012", "failed to get attendances", http.StatusInternalServerError)
	// CoreFailedToUpdateAttendances failed to update attendances
	CoreFailedToUpdateAttendances = new(core, "E2013", "failed to update attendances", http.StatusInternalServerError)
	// CoreFailedToGetPhysiques failed to get physiques
	CoreFailedToGetPhysiques = new(core, "E2014", "failed to get physiques", http.StatusInternalServerError)
	// CoreFailedToUpdatePhysiques failed to update physiques
	CoreFailedToUpdatePhysiques = new(core, "E2015", "failed to update physiques", http.StatusInternalServerError)
	// CoreInvalidPupil invalid pupil
	CoreInvalidPupil = new(core, "E2016", "invalid pupil", http.StatusBadRequest)
	// CoreInvalidPhysique invalid physique (failed to resolve)
	CoreInvalidPhysique = new(core, "E2017", "invalid physique", http.StatusBadRequest)
	// CoreFailedToGetPhysiqueMasters failed to get physique masters
	CoreFailedToGetPhysiqueMasters = new(core, "E2018", "failed to get physique masters", http.StatusInternalServerError)
	// CoreFailedToGetIngredient failed to get ingredient
	CoreFailedToGetIngredient = new(core, "E2019", "failed to get ingredient", http.StatusInternalServerError)
	// CoreFailedToSaveIngredient failed to save ingredient
	CoreFailedToSaveIngredient = new(core, "E2020", "failed to save ingredient", http.StatusInternalServerError)
	// CoreFailedToSaveRecipe failed to save recipe
	CoreFailedToSaveRecipe = new(core, "E2021", "failed to save recipe", http.StatusInternalServerError)
	// CoreFailedToGetMenu failed to get menu
	CoreFailedToGetMenu = new(core, "E2022", "failed to get menu", http.StatusInternalServerError)
	// CoreFailedToGetRecipe failed to get recipe
	CoreFailedToGetRecipe = new(core, "E2023", "failed to get recipe", http.StatusInternalServerError)
	// CoreFailedToGetProcurement failed to get procurements
	CoreFailedToGetProcurement = new(core, "E2024", "failed to get procurements", http.StatusInternalServerError)
	// CoreInvalidAttendanceCountRequest invalid attendance count request
	CoreInvalidAttendanceCountRequest = new(core, "E2025", "invalid attendance count request", http.StatusBadRequest)
	// CoreFailedToGetAttendanceCount invalid attendance count request
	CoreFailedToGetAttendanceCount = new(core, "E2026", "failed to get attendance count", http.StatusInternalServerError)
	// CoreInvalidIngredientCategory invalid ingredient category
	CoreInvalidIngredientCategory = new(core, "E2027", "invalid ingredient category", http.StatusBadRequest)
	// CoreFailedToGetGrowthProfile failed to get growth profile
	CoreFailedToGetGrowthProfile = new(core, "E2028", "failed to get growth profile", http.StatusInternalServerError)
	// CoreFailedToSaveGrowthProfile failed to save growth profile
	CoreFailedToSaveGrowthProfile = new(core, "E2029", "failed to save growth profile", http.StatusInternalServerError)
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
