package filter

import (
	"errors"

	"github.com/disintegration/imaging"
)

type Filter interface {
	Process(src, dst, name string) error
	SetSigma(sigma float64)
}

type FilterOpt string

var filterOpt = map[string]bool{
	"grayscale": true,
	"blur":      true,
}

type FilterImg struct {
	filter string
	sigma  float64
}

func New(filter string) (Filter, error) {
	if !filterOpt[filter] {
		return nil, errors.New("wrong filter option")
	}
	obj := &FilterImg{filter: filter}
	return obj, nil

}

func (f FilterImg) Process(src, dst, name string) error {
	in := src + "/" + name
	switch f.filter {
	case "grayscale":
		pathOut := dst + "/" + "grayscale_" + name
		return f.processGrayscale(in, pathOut)
	case "blur":
		pathOut := dst + "/" + "blur" + name
		return f.processBlur(in, pathOut)
	}
	return nil
}

func (f FilterImg) processGrayscale(name, pathOut string) error {
	img, err := imaging.Open(name)
	if err != nil {
		return err
	}
	imgGS := imaging.Grayscale(img)
	return imaging.Save(imgGS, pathOut)
}

func (f FilterImg) processBlur(name, pathOut string) error {
	img, err := imaging.Open(name)
	if err != nil {
		return err
	}
	imgGS := imaging.Blur(img, f.sigma)
	return imaging.Save(imgGS, pathOut)
}

func (f *FilterImg) SetSigma(sigma float64) {
	f.sigma = sigma
}
