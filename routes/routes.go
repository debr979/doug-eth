package routes

import (
	"doug/controllers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type routes struct{}

var Routes routes

func (r *routes) Server() error {

	gin.SetMode(os.Getenv("RUN_MODE"))
	eng := gin.Default()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.routes(eng)

	log.Printf("Run service on port %s ...", port)
	return eng.Run(port)
}

func (r *routes) routes(e *gin.Engine) {
	//Blocks
	e.GET("/blocks", controllers.Blocks.GetBlockByCount)
	e.GET("/blocks/:id", controllers.Blocks.GetBlockById)

	//Transactions
	e.GET("/transaction/:txHash", controllers.Transactions.GetTransactionByHash)
}
