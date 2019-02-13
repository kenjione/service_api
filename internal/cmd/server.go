package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kenjione/importer"
	"github.com/kenjione/service_api/internal/app"
	"github.com/kenjione/service_api/internal/config"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts BE server",
	Run: func(cmd *cobra.Command, args []string) {
		cnf := makeConfigFromConfigFlag(cmd)
		srv := MakeServer(cnf)

		defer srv.Importer.Close()
		srv.App.Run(cnf.ServerPort)
	},
}

func Execute() {
	if err := serverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func MakeServer(cnf *config.ServerConfig) *app.LocationServer {
	importerService := importer.NewImporter(&importer.Config{
		DatabaseName:     cnf.DatabaseName,
		DatabaseUser:     cnf.DatabaseUser,
		DatabasePassword: cnf.DatabasePassword,
		DatabaseAddr:     cnf.DatabaseAddr,
	})

	server := &app.LocationServer{}
	ginApp := gin.Default()

	ginApp.GET("/locations/:ip", server.LocationHandler)

	server.App = ginApp
	server.Importer = importerService

	return server
}

func init() {
	serverCmd.Flags().String("config", "configs/config.json", "Config file to load")
}

func makeConfigFromConfigFlag(cmd *cobra.Command) *config.ServerConfig {
	configFile, err := filepath.Abs(cmd.Flag("config").Value.String())
	if err != nil {
		log.Fatal("can not get full path to config file:", err)
	}

	cnf, err := config.MakeServerConfigFromFile(configFile)
	if err != nil {
		log.Fatal("error loading config: ", err)
	}
	return cnf
}
