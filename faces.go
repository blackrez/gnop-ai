package main

import (
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/owulveryck/gofaces"
	"github.com/owulveryck/gofaces/draw"
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
	"gorgonia.org/tensor"
)

type configuration struct {
	ConfidenceThreshold float64 `envconfig:"confidence_threshold" default:"0.10" required:"true"`
	ClassProbaThreshold float64 `envconfig:"proba_threshold" default:"0.90" required:"true"`
}

var (
	model   = flag.String("model", "model/model.onnx", "path to the model file")
	debug  = flag.Bool("s", false, "silent mode (useful if output is -)")
	config  configuration
)

func main() {
		router.Static("/", "public")

	router.POST("/upload", func(c *gin.Context) {

		imgF, _ = ioutil.TempFile("", "int-*.jpeg")
		defer os.Remove(imgF.Name())
		outputF, _ = ioutil.TempFile("", "out-*.png")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		if err := c.SaveUploadedFile(file, imgF.Name()); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		buf, _ := ioutil.ReadFile(imgF.Name())
		if !filetype.IsImage(buf) {
			log.Println("Not an image")
			c.JSON(http.StatusBadRequest, gin.H{"result": "file is not a image"})
			return
		}
		// Decode it into the model
	// Create a backend receiver
	backend := gorgonnx.NewGraph()
	// Create a model and set the execution backend
	m := onnx.NewModel(backend)

	// read the onnx model
	b, err := ioutil.ReadFile(*model)
	if err != nil {
		log.Fatal(err)
	}
	// Decode it into the model
	err = m.UnmarshalBinary(b)
	if err != nil {
		log.Fatal(err)
	}

	img, err := os.Open(*imgF)
	if err != nil {
		log.Fatal(err)
	}
	inputT, err := gofaces.GetTensorFromImage(img)
	if err != nil {
		log.Fatal(err)
	}
	m.SetInput(0, inputT)
	err = backend.Run()
	if err != nil {
		log.Fatal(err)
	}
	outputs, err := m.GetOutputTensors()

	boxes, err := gofaces.ProcessOutput(outputs[0].(*tensor.Dense))
	if err != nil {
		log.Fatal(err)
	}
	boxes = gofaces.Sanitize(boxes)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(boxes); i++ {
		if boxes[i].Confidence < config.ConfidenceThreshold {
			boxes = boxes[:i]
			//continue
		}
		/*
			if boxes[i].Classes[0].Prob < config.ClassProbaThreshold {
				boxes = boxes[:i]
				continue
			}
		*/
	}

	fmt.Println(boxes)

	if *outputF != "" {
		mask := draw.CreateMask(gofaces.WSize, gofaces.HSize, boxes)
		f, err := os.Create(*outputF)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		err = png.Encode(f, mask)
		if err != nil {
			log.Fatal(err)
		}
	}

		c.JSON(http.StatusOK, gin.H{"result": result})
	})
	router.Run(":8080")
}