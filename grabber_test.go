package grabber

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestResolveAddr(t *testing.T) {
	ta := assert.New(t)
	ta.Equal([]string{"http://www.example.com/img_033.jpg"}, ResolveAddr("http://www.example.com/img_033.jpg"))
	ta.Equal([]string{
		"http://www.example.com/img_8.jpg",
		"http://www.example.com/img_9.jpg",
		"http://www.example.com/img_10.jpg",
		"http://www.example.com/img_11.jpg",
		"http://www.example.com/img_12.jpg",
		}, ResolveAddr("http://www.example.com/img_[8:12].jpg"))
	ta.Equal([]string{
		"http://www.example.com/img_08.jpg",
		"http://www.example.com/img_09.jpg",
		"http://www.example.com/img_10.jpg",
		"http://www.example.com/img_11.jpg",
		"http://www.example.com/img_12.jpg",
	}, ResolveAddr("http://www.example.com/img_[08:12].jpg"))
	ta.Equal([]string{"http://www.example.com/img_[008:12].jpg"}, ResolveAddr("http://www.example.com/img_[008:12].jpg"))
}
