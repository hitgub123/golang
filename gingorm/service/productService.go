package service
import(
	"slq.me/dao"
	"slq.me/obj"
	// "fmt"
)

func PageProduct() obj.Result{
	res:=obj.Result{Code: 0,Error: nil,Data: "page"}
	return res

}

func OneProduct(id string) obj.Result{
	dao.OneProduct(id)
	res:=obj.Result{Code: 0,Error: nil,Data: "page"}
	return res

}

func DeleteProduct(id string) obj.Result{
	dao.DeleteProduct(id)
	res:=obj.Result{Code: 0,Error: nil,Data: "page"}
	return res
}

func UpdateProduct(o obj.User) obj.Result{
	dao.UpdateProduct(o)
	res:=obj.Result{Code: 0,Error: nil,Data: "page"}
	return res
}

func InsertProduct(o obj.User) obj.Result{
	dao.InsertProduct(o)
	res:=obj.Result{Code: 0,Error: nil,Data: "page"}
	return res
}