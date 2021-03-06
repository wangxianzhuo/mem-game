package pic

// Pic 图片资源
type Pic struct {
	ID  int    `json:"id"`
	URI string `json:"uri"`
}

// LoadPics 加载图片资源
func LoadPics(path string, num int) ([]Pic, error) {
	pics := make([]Pic, num)
	for index := 0; index < num; index++ {
		pics[index].ID = index
	}
	return pics, nil
}
