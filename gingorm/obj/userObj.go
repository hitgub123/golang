package obj
import(
    "gorm.io/gorm"
    "time"
)
type User struct {
    gorm.Model
	Id        int       `json:"id" form:"id" gorm:"column:id;unique;primaryKey;autoIncrement"`
    Name  string	    `json:"name" form:"name"  gorm:"column:name"`
	Age       int       `json:"age" form:"age"  gorm:"column:age"`
    Password string		`json:"password" form:"password"  gorm:"column:password"`
	Lasttime  time.Time `json:"lasttime" form:"lasttime"  gorm:"column:lasttime"`
}

//定义表名
func (User) TableName() string {
	return "users"
}
