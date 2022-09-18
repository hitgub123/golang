package controller
import(
	// "github.com/gin-gonic/gin"
	"slq.me/service"
	"slq.me/obj"
	"fmt"
)

func GetProduct(){
	service.PageProduct()
	Product:=obj.Product{}
	fmt.Printf("%v", Product)	
}