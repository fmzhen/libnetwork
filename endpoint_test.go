package libnetwork

import (
	"os"
	"testing"
)

func TestCreateBasePath(t *testing.T) {

	//a not exists path, there non't check the permit
	basePath := "/home/fmzhen/libgotest"

	if err := createBasePath(basePath); err != nil {
		t.Errorf("%v", err)
	}

	//stat()ruturn FileInfo
	stat, err := os.Stat(basePath)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !stat.IsDir() {
		t.Errorf("createBasePath Fail")
	}

	if stat.Mode() != 0644 {
		t.Errorf("createBasePath Fail,the FileMode error")
	}
}
