package data

import(
	"os"
	"time"
)


func dictUnidictBnBytes() ([]byte, error) {
	return _dictUnidictBn, nil
}

func dictUnidictBn() (*asset, error) {
	bytes, err := dictUnidictBnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
