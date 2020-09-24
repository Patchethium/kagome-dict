package data

import(
	"os"
	"time"
)


func dictIpadictAnBytes() ([]byte, error) {
	return _dictIpadictAn, nil
}

func dictIpadictAn() (*asset, error) {
	bytes, err := dictIpadictAnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.an", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
