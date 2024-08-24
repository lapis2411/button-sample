package resource

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"lapis2411/button-sample/types"
)

func createBoxImage(text string, size types.Size, col color.RGBA) *image.RGBA {

	// 薄い青の矩形を描画するために画像を生成
	img := image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height)))

	// 背景色（薄い青）を設定
	draw.Draw(img, img.Bounds(), &image.Uniform{col}, image.Point{}, draw.Src)

	// 文字 "Start" を中央に描画する
	fcolor := color.Black
	point := fixed.Point26_6{
		X: fixed.I(20),                   // 文字の幅に合わせて調整
		Y: fixed.I(int(size.Height) / 2), // 文字の高さに合わせて調整
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(fcolor),
		Face: basicfont.Face7x13, // 簡単なフォント
		Dot:  point,
	}
	d.DrawString(text)
	return img
}
