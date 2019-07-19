package distribution

/**
cd framework/distribution
go test
*/
import (
	"testing"
	//"os"
)

func Test_DataID(t *testing.T) {
	mid  := DataID(4,1)
	println("[Test] get value", mid)
	if mid  > 0 {
		t.Log("Test_DataID 通过")
	}
}

// func TestMain(m *testing.M) {
// 	println("[Test] test begin")
// 	execPath, err := os.Executable()
// 	pwd, err := os.Getwd()
// 	println("[Test] pwd:%s",pwd)
// 	println("[Test] execPath:%s",execPath)
// 	err = os.Setenv("IS_DEBUG", "true")
// 	if err != nil {
// 		println(err)
// 	}
// 	m.Run()
// 	println("[Test] test end")
// }
