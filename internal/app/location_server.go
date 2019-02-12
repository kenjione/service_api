package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kenjione/importer"
)

type LocationServer struct {
	Importer *importer.Importer
	App      *gin.Engine
}
