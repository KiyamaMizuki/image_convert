# image_convert
png型式からjpegに変換するプログラムです。
pathを指定すると再起的に探索して指定したディレクトリ配下のpng画像をjpgに変換し元のpng画像を消します。
※注意絶対パスで指定しないとエラーを起こします

指定の方法としては以下の通りです  
```go run convert_png_to_jpg.go < pass```
