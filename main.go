package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gopkg.in/h2non/bimg.v1"
)

func mnistHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func convertImg(){
buffer, err := bimg.Read("image.jpg")
if err != nil {
  fmt.Fprintln(os.Stderr, err)
}

newImage, err := bimg.NewImage(buffer).Colourspace(bimg.INTERPRETATION_B_W)
if err != nil {
  fmt.Fprintln(os.Stderr, err)
}

colourSpace, _ := bimg.ImageInterpretation(newImage)
if colourSpace != bimg.INTERPRETATION_B_W {
  fmt.Fprintln(os.Stderr, "Invalid colour space")
}
}

func main() {
    http.HandleFunc("/mnist", mnistHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
