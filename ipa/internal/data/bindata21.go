package data

import(
	"os"
	"time"
)


func dictIpadictAvBytes() ([]byte, error) {
	return _dictIpadictAv, nil
}

func dictIpadictAv() (*asset, error) {
	bytes, err := dictIpadictAvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.av", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596545184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
