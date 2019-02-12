package app

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/kenjione/service_api/internal"
)

func (s *LocationServer) LocationHandler(c *gin.Context) {
	ipAddress := c.Query("ip_address")
	lbytes, err := s.Importer.FindByIP(ipAddress)

	if err != nil && err.Error() == "pg: no rows in result set" {
		c.JSON(404, gin.H{"response": "Location not found"})
		return
	}

	location := &internal.Location{}
	err = json.Unmarshal(lbytes, location)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"response": location})
	}
}
