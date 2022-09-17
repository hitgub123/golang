package main
import(
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if ex,err:=os.Executable();err==nil{
		//貌似没啥用
		//C:\Users\81802\AppData\Local\Temp\go-build2281919180\b001\exe
		fmt.Printf(filepath.Dir(ex))
		
		//C:\Users\81802\AppData\Local\Temp\go-build1409437024\b001\exe\u1.exe
		// fmt.Printf("%v", ex)
	}
	
}