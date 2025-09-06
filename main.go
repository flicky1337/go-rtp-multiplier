package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	rtp := flag.Float64("rtp", 0.95, "target RTP value between (0, 1.0]")
	flag.Parse()
	if *rtp <= 0 || *rtp > 1 {
		log.Fatalf("invalid RTP value: %f (must be in (0,1])", *rtp)
	}

	r := gin.Default()
	r.GET("/get", getHandler(*rtp))
	r.Run(":64333")
}
