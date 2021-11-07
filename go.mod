module kreoapp2

go 1.17

require (
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/gorilla/websocket v1.4.2
	github.com/grokify/html-strip-tags-go v0.0.1
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/russross/blackfriday v1.6.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/satori/go.uuid v1.2.0 => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
