package main

import(
    "os"
    "fmt"
    "bufio"
    "strings"
    "io/ioutil"
    "path/filepath"
    "image/jpeg"
    "image/png"
    
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    var name string
    fmt.Scan(&name)
    //ファイルオープン
    var paths []string = getCurrentDirectory(name)
    for _, v := range paths{
        if(strings.HasSuffix(v, "png")){
            file, err := os.Open(v)
            assert(err, "Invalid image file path" + v)
            defer file.Close()

            // ファイルオブジェクトを画像オブジェクトに変換
            fmt.Println(v)
            img, err := png.Decode(file)
            assert(err, "imagejjj")

            // 出力画像パス.
            dstPath := strings.Replace(v, "png", "jpg", 1)

            // 出力ファイルを生成
            out, err := os.Create(dstPath)
            defer out.Close()

            // 画像ファイル出力
            jpeg.Encode(out, img, nil)

            //最後にpngを消す
            os.Remove(v);
        }
        
    }
}

// errorオブジェクトをチェックし、nilの場合例外を送出
func assert(err error, msg string) {
    if err != nil {
        panic(err.Error() + ":" + msg)
    }
}

//ファイルまでの相対パス取得
func getCurrentDirectory(dir string) []string{
    //dir := "."
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

    var paths []string
    for _, file := range files {
        if file.IsDir() {
            paths = append(paths, getCurrentDirectory(filepath.Join(dir, file.Name()))...)
            continue
        }
        paths = append(paths, filepath.Join(dir, file.Name()))
    }
    return paths
}