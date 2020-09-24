package data

import(
	"os"
	"time"
)


func dictUnidictBeBytes() ([]byte, error) {
	return _dictUnidictBe, nil
}

func dictUnidictBe() (*asset, error) {
	bytes, err := dictUnidictBeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.be", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1596854809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
