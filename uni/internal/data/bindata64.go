package data

import(
	"os"
	"time"
)


func dictUnidictCmBytes() ([]byte, error) {
	return _dictUnidictCm, nil
}

func dictUnidictCm() (*asset, error) {
	bytes, err := dictUnidictCmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
