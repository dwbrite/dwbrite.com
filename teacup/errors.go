package teacup

import (
	"html/template"
	"log"
	"net/http"
)

type errorPage struct {
	Status int
	Message string
}

type tcErrors struct {
	template *template.Template
	errorMessages map[int]string
}

func newTcErrors() tcErrors {
	return tcErrors{
		template.Must(template.New("tmpl") .ParseFiles("teacup/error_template.gohtml")),
		map[int]string{
			http.StatusBadRequest:                   "Bad Request",
			http.StatusUnauthorized:                 "Unauthorized",
			http.StatusPaymentRequired:              "Payment Required",
			http.StatusForbidden:                    "Forbidden",
			http.StatusNotFound:                     "Not Found",
			http.StatusMethodNotAllowed:             "Method Not Allowed",
			http.StatusNotAcceptable:                "Not Acceptable",
			http.StatusProxyAuthRequired:            "Proxy Authentication Required",
			http.StatusRequestTimeout:               "Request Timeout",
			http.StatusConflict:                     "Conflict",
			http.StatusGone:                         "Gone",
			http.StatusLengthRequired:               "Length Required",
			http.StatusPreconditionFailed:           "Precondition Failed",
			http.StatusRequestEntityTooLarge:        "Request Entity Too Large",
			http.StatusRequestURITooLong:            "Request URI Too Long",
			http.StatusUnsupportedMediaType:         "Unsupported Media Type",
			http.StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
			http.StatusExpectationFailed:            "Expectation Failed",
			http.StatusTeapot:                       "I'm a teapot",
			http.StatusMisdirectedRequest:           "Misdirected Request",
			http.StatusUnprocessableEntity:          "Unprocessable Entity",
			http.StatusLocked:                       "Locked",
			http.StatusFailedDependency:             "Failed Dependency",
			http.StatusUpgradeRequired:              "Upgrade Required",
			http.StatusPreconditionRequired:         "Precondition Required",
			http.StatusTooManyRequests:              "Too Many Requests",
			http.StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
			http.StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

			http.StatusInternalServerError:           "Internal Server Error",
			http.StatusNotImplemented:                "Not Implemented",
			http.StatusBadGateway:                    "Bad Gateway",
			http.StatusServiceUnavailable:            "Service Unavailable",
			http.StatusGatewayTimeout:                "Gateway Timeout",
			http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
			http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
			http.StatusInsufficientStorage:           "Insufficient Storage",
			http.StatusLoopDetected:                  "Loop Detected",
			http.StatusNotExtended:                   "Not Extended",
			http.StatusNetworkAuthenticationRequired: "Network Authentication Required",
		},
	}
}

func (e *tcErrors) setTemplate(filename string) {
	e.template = template.Must(template.ParseFiles(filename))
}

func (e *tcErrors) setErrorText(code int, media string) {
	e.errorMessages[code] = media
}

func (t *teacup) serveError(writer http.ResponseWriter, httpStatus int) {
	writer.WriteHeader(httpStatus)
	t.errors.template.Execute(writer,
		errorPage{
			httpStatus,
			t.errors.errorMessages[httpStatus],
	})
}

func (t teacup) checkAndLogError(err error) bool {
	isError := err != nil
	if isError {
		log.Println(err)
		t.Log.Println(err)
	}
	return isError
}