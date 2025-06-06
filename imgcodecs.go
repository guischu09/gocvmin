package gocv

/*
#include <stdlib.h>
#include "imgcodecs.h"
*/
import "C"
import (
	"unsafe"
)

// IMReadFlag is one of the valid flags to use for the IMRead function.
type IMReadFlag int

const (
	// IMReadUnchanged return the loaded image as is (with alpha channel,
	// otherwise it gets cropped).
	IMReadUnchanged IMReadFlag = -1

	// IMReadGrayScale always convert image to the single channel
	// grayscale image.
	IMReadGrayScale IMReadFlag = 0

	// IMReadColor always converts image to the 3 channel BGR color image.
	IMReadColor IMReadFlag = 1

	// IMReadAnyDepth returns 16-bit/32-bit image when the input has the corresponding
	// depth, otherwise convert it to 8-bit.
	IMReadAnyDepth IMReadFlag = 2

	// IMReadAnyColor the image is read in any possible color format.
	IMReadAnyColor IMReadFlag = 4

	// IMReadLoadGDAL uses the gdal driver for loading the image.
	IMReadLoadGDAL IMReadFlag = 8

	// IMReadReducedGrayscale2 always converts image to the single channel grayscale image
	// and the image size reduced 1/2.
	IMReadReducedGrayscale2 IMReadFlag = 16

	// IMReadReducedColor2 always converts image to the 3 channel BGR color image and the
	// image size reduced 1/2.
	IMReadReducedColor2 IMReadFlag = 17

	// IMReadReducedGrayscale4 always converts image to the single channel grayscale image and
	// the image size reduced 1/4.
	IMReadReducedGrayscale4 IMReadFlag = 32

	// IMReadReducedColor4 always converts image to the 3 channel BGR color image and
	// the image size reduced 1/4.
	IMReadReducedColor4 IMReadFlag = 33

	// IMReadReducedGrayscale8 always convert image to the single channel grayscale image and
	// the image size reduced 1/8.
	IMReadReducedGrayscale8 IMReadFlag = 64

	// IMReadReducedColor8 always convert image to the 3 channel BGR color image and the
	// image size reduced 1/8.
	IMReadReducedColor8 IMReadFlag = 65

	// IMReadIgnoreOrientation do not rotate the image according to EXIF's orientation flag.
	IMReadIgnoreOrientation IMReadFlag = 128
)

// TODO: Define IMWriteFlag type?

const (
	//IMWriteJpegQuality is the quality from 0 to 100 for JPEG (the higher is the better). Default value is 95.
	IMWriteJpegQuality = 1

	// IMWriteJpegProgressive enables JPEG progressive feature, 0 or 1, default is False.
	IMWriteJpegProgressive = 2

	// IMWriteJpegOptimize enables JPEG optimization, 0 or 1, default is False.
	IMWriteJpegOptimize = 3

	// IMWriteJpegRstInterval is the JPEG restart interval, 0 - 65535, default is 0 - no restart.
	IMWriteJpegRstInterval = 4

	// IMWriteJpegLumaQuality separates luma quality level, 0 - 100, default is 0 - don't use.
	IMWriteJpegLumaQuality = 5

	// IMWriteJpegChromaQuality separates chroma quality level, 0 - 100, default is 0 - don't use.
	IMWriteJpegChromaQuality = 6

	// IMWritePngCompression is the compression level from 0 to 9 for PNG. A
	// higher value means a smaller size and longer compression time.
	// If specified, strategy is changed to IMWRITE_PNG_STRATEGY_DEFAULT (Z_DEFAULT_STRATEGY).
	// Default value is 1 (best speed setting).
	IMWritePngCompression = 16

	// IMWritePngStrategy is one of cv::IMWritePNGFlags, default is IMWRITE_PNG_STRATEGY_RLE.
	IMWritePngStrategy = 17

	// IMWritePngBilevel is the binary level PNG, 0 or 1, default is 0.
	IMWritePngBilevel = 18

	// IMWritePxmBinary for PPM, PGM, or PBM can be a binary format flag, 0 or 1. Default value is 1.
	IMWritePxmBinary = 32

	// IMWriteWebpQuality is the quality from 1 to 100 for WEBP (the higher is
	// the better). By default (without any parameter) and for quality above
	// 100 the lossless compression is used.
	IMWriteWebpQuality = 64

	// IMWritePamTupletype sets the TUPLETYPE field to the corresponding string
	// value that is defined for the format.
	IMWritePamTupletype = 128

	// IMWritePngStrategyDefault is the value to use for normal data.
	IMWritePngStrategyDefault = 0

	// IMWritePngStrategyFiltered is the value to use for data produced by a
	// filter (or predictor). Filtered data consists mostly of small values
	// with a somewhat random distribution. In this case, the compression
	// algorithm is tuned to compress them better.
	IMWritePngStrategyFiltered = 1

	// IMWritePngStrategyHuffmanOnly forces Huffman encoding only (no string match).
	IMWritePngStrategyHuffmanOnly = 2

	// IMWritePngStrategyRle is the value to use to limit match distances to
	// one (run-length encoding).
	IMWritePngStrategyRle = 3

	// IMWritePngStrategyFixed is the value to prevent the use of dynamic
	// Huffman codes, allowing for a simpler decoder for special applications.
	IMWritePngStrategyFixed = 4
)

// IMRead reads an image from a file into a Mat.
// The flags param is one of the IMReadFlag flags.
// If the image cannot be read (because of missing file, improper permissions,
// unsupported or invalid format), the function returns an empty Mat.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga288b8b3da0892bd651fce07b3bbd3a56
func IMRead(name string, flags IMReadFlag) Mat {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return newMat(C.Image_IMRead(cName, C.int(flags)))
}

// IMReadMulti reads multi-page image from a file into a []Mat.
// The flags param is one of the IMReadFlag flags.
// If the image cannot be read (because of missing file, improper permissions,
// unsupported or invalid format), the function returns an empty []Mat.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d4/da8/group__imgcodecs.html#gaaeb5e219c706fd6aa1ec6cf17b172080
func IMReadMulti(name string, flags IMReadFlag) []Mat {
	var mats []Mat
	multiRead := C.struct_Mats{}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	multiRead = C.Image_IMReadMulti(cName, C.int(flags))
	defer C.Mats_Close(multiRead)

	if multiRead.length > C.int(0) {
		mats = make([]Mat, multiRead.length)
		for i := 0; i < int(multiRead.length); i++ {
			mats[i].p = C.Mats_get(multiRead, C.int(i))
		}
	}
	return mats
}

// IMReadMulti reads multi-page image from a file into a []Mat.
// The flags param is one of the IMReadFlag flags.
// If the image cannot be read (because of missing file, improper permissions,
// unsupported or invalid format), the function returns an empty []Mat.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d4/da8/group__imgcodecs.html#ga55e88dc40b65807cfbe2c62d27f7fdf9
func IMReadMulti_WithParams(name string, start int, count int, flags IMReadFlag) []Mat {
	var mats []Mat
	multiRead := C.struct_Mats{}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	multiRead = C.Image_IMReadMulti_WithParams(cName, C.int(start), C.int(count), C.int(flags))
	defer C.Mats_Close(multiRead)

	if multiRead.length > C.int(0) {
		mats = make([]Mat, multiRead.length)
		for i := 0; i < int(multiRead.length); i++ {
			mats[i].p = C.Mats_get(multiRead, C.int(i))
		}
	}
	return mats
}

// IMWrite writes a Mat to an image file.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#gabbc7ef1aa2edfaa87772f1202d67e0ce
func IMWrite(name string, img Mat) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return bool(C.Image_IMWrite(cName, img.p))
}

// IMWriteWithParams writes a Mat to an image file. With that func you can
// pass compression parameters.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#gabbc7ef1aa2edfaa87772f1202d67e0ce
func IMWriteWithParams(name string, img Mat, params []int) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cparams := []C.int{}

	for _, v := range params {
		cparams = append(cparams, C.int(v))
	}

	paramsVector := C.struct_IntVector{}
	paramsVector.val = (*C.int)(&cparams[0])
	paramsVector.length = (C.int)(len(cparams))

	return bool(C.Image_IMWrite_WithParams(cName, img.p, paramsVector))
}

// FileExt represents a file extension.
type FileExt string

const (
	// PNGFileExt is the file extension for PNG.
	PNGFileExt FileExt = ".png"
	// JPEGFileExt is the file extension for JPEG.
	JPEGFileExt FileExt = ".jpg"
	// GIFFileExt is the file extension for GIF.
	GIFFileExt FileExt = ".gif"
)

// IMEncode encodes an image Mat into a memory buffer.
// This function compresses the image and stores it in the returned memory buffer,
// using the image format passed in in the form of a file extension string.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga461f9ac09887e47797a54567df3b8b63
func IMEncode(fileExt FileExt, img Mat) (buf *NativeByteBuffer, err error) {
	cfileExt := C.CString(string(fileExt))
	defer C.free(unsafe.Pointer(cfileExt))

	buffer := newNativeByteBuffer()
	res := C.Image_IMEncode(cfileExt, img.Ptr(), buffer.nativePointer())
	return buffer, OpenCVResult(res)
}

// IMEncodeWithParams encodes an image Mat into a memory buffer.
// This function compresses the image and stores it in the returned memory buffer,
// using the image format passed in in the form of a file extension string.
//
// Usage example:
//
//	buffer, err := gocv.IMEncodeWithParams(gocv.JPEGFileExt, img, []int{gocv.IMWriteJpegQuality, quality})
//
// For further details, please see:
// http://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga461f9ac09887e47797a54567df3b8b63
func IMEncodeWithParams(fileExt FileExt, img Mat, params []int) (buf *NativeByteBuffer, err error) {
	cfileExt := C.CString(string(fileExt))
	defer C.free(unsafe.Pointer(cfileExt))

	cparams := []C.int{}

	for _, v := range params {
		cparams = append(cparams, C.int(v))
	}

	paramsVector := C.struct_IntVector{}
	paramsVector.val = (*C.int)(&cparams[0])
	paramsVector.length = (C.int)(len(cparams))

	b := newNativeByteBuffer()
	res := C.Image_IMEncode_WithParams(cfileExt, img.Ptr(), paramsVector, b.nativePointer())
	return b, OpenCVResult(res)
}

// IMDecode reads an image from a buffer in memory.
// The function IMDecode reads an image from the specified buffer in memory.
// If the buffer is too short or contains invalid data, the function
// returns an empty matrix.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga26a67788faa58ade337f8d28ba0eb19e
func IMDecode(buf []byte, flags IMReadFlag) (Mat, error) {
	data, err := toByteArray(buf)
	if err != nil {
		return Mat{}, err
	}
	return newMat(C.Image_IMDecode(*data, C.int(flags))), LastExceptionError()
}

// IMDecodeIntoMat reads an image from a buffer in memory into a matrix.
// The function IMDecodeIntoMat reads an image from the specified buffer in memory.
// If the buffer is too short or contains invalid data, the function
// returns an error
//
// For further details, please see:
// https://docs.opencv.org/4.x/d4/da8/group__imgcodecs.html#ga5a0acefe5cbe0a81e904e452ec7ca733
func IMDecodeIntoMat(buf []byte, flags IMReadFlag, dest *Mat) error {
	data, err := toByteArray(buf)
	if err != nil {
		return err
	}
	return OpenCVResult(C.Image_IMDecodeIntoMat(*data, C.int(flags), dest.p))
}
