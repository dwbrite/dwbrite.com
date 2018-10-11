package teacup

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type teacup struct {
	Port          	uint16
	DbInfo        	string
	TlsPair       	*TlsKeyPair

	FileWhitelist 	regexp.Regexp
	DirBlacklist  	regexp.Regexp

	Log           	*log.Logger
	errors     		tcErrors

	tables 		  	map[string]bool
	webpages		map[string]dynamicPage
}

type dynamicPage struct {
	template 	*template.Template
	table 		string
	fn			func(http.Request, string) interface{}
}

type TlsKeyPair struct {
	Cert string
	Key  string
}

type PageContents struct {
	Uid      uint32
	Summary  sql.NullString
	Title    string
	Body     template.HTML
	PostDate time.Time
}


func NewTeacup(port uint16, dbInfo string, pair *TlsKeyPair,  fileWhitelist regexp.Regexp, dirBlacklist regexp.Regexp, log *log.Logger) *teacup {
	t := teacup {
		port,
		dbInfo,
		pair,
		fileWhitelist,
		dirBlacklist,
		log,
		newTcErrors(),
		make(map[string]bool),
		make(map[string]dynamicPage),

	}
	return &t
}

func (t *teacup) StartServer() {
	http.HandleFunc("/", t.matchRequest)
	if t.TlsPair == nil {
		t.Log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(t.Port)), nil))
	} else {
		t.Log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(int(t.Port)), t.TlsPair.Cert, t.TlsPair.Key, nil))
	}
}

func (t *teacup) matchRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	switch {
	case t.DirBlacklist.MatchString(path):
		t.serveError(writer, http.StatusForbidden)
	default:
		t.servePage(writer, request)
	}
}

func (t *teacup) serveFile(writer http.ResponseWriter, request *http.Request) {
	if t.FileWhitelist.MatchString(request.URL.Path) {
		http.ServeFile(writer, request, request.URL.Path[1:])
	} else {
		t.serveError(writer, http.StatusNotFound)
	}
}

func (t *teacup) servePage(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	for pathRegex := range t.webpages {
		if regexp.MustCompile(pathRegex).MatchString(path) {
			dynPage := t.webpages[pathRegex]
			content := dynPage.fn(*request, t.DbInfo)
			if content == nil { break }

			err := dynPage.template.Execute(writer, content)
			if t.checkAndLogError(err) {
				t.serveError(writer, http.StatusBadRequest)
			}
			return
		}
	}

	t.serveFile(writer, request)
}

func (t *teacup) AddDynamicPage(pathRegex string, table string, tmplFnMap template.FuncMap, fn func(http.Request, string) interface{}, templates ... string) {
	regexp.MustCompile(pathRegex)
	tmpl := template.Must(template.New("teacup_base").Funcs(tmplFnMap).ParseFiles(templates...))
	if strings.Contains(tmpl.DefinedTemplates(), `"teacup_base"`) {
		t.webpages[pathRegex] = dynamicPage{tmpl, table, fn}
		return
	}
	t.Log.Panicln(`Template error: Please define a template with the name "teacup_base"`)
}