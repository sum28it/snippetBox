package main

import (
	"github.com/alexedwards/scs"
	"github.com/sum28it/snippetBox/pkg/models"
)

type App struct {
	Addr      string
	Database  *models.Database
	HTMLDir   string
	Sessions  *scs.Manager
	StaticDir string
	TLSCert   string
	TLSKey    string
}
