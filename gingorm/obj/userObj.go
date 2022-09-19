package obj
import(
    "gorm.io/gorm"
    "time"
)
type User struct {
    gorm.Model
	Id        int       `json:"id" form:"id" gorm:"column:id;unique;primaryKey;autoIncrement"`
    Name  string	    `json:"name" form:"name" binding:"required,min=2,max=4" gorm:"column:name"`
	Age       int       `json:"age" form:"age" binding:"required,gt=2,lte=4" gorm:"column:age"`
    Password string		`json:"password" form:"password"  gorm:"column:password"`
	Lasttime  time.Time `json:"lasttime" form:"lasttime"  gorm:"column:lasttime"`
    Pic1 string          `json:"pic1" form:"pic1"  gorm:"column:pic1`
    Pic2 string          `json:"pic2" form:"pic2"  gorm:"column:pic2`
}

//定义表名
func (User) TableName() string {
	return "users"
}
