package dao
import(
	"slq.me/obj"
	"fmt"

)

func SelectProduct()obj.Pagination{
	product:= obj.Product{

	}
	fmt.Printf("%v", product)
	o:= obj.Pagination{}
	return o
}

func OneProduct(id string)obj.User{
	o:= obj.User{}
	return o

}

func DeleteProduct(id string){
	

}

func UpdateProduct(o obj.User){


}


func InsertProduct(o obj.User){


}