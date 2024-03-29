package app

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/kenjione/service_api/internal"
)

func (s *LocationServer) LocationHandler(c *gin.Context) {
	ipAddress := c.Param("ip")
	lbytes, err := s.Importer.FindByIP(ipAddress)

	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			c.JSON(404, gin.H{"response": "Location not found"})
			return
		}

		c.JSON(500, gin.H{"response": err.Error()})
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
