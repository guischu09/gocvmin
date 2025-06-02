//go:build !customenv && !opencvstatic

package gocv

// Changes here should be mirrored in contrib/cgo.go and cuda/cgo.go.

/*
#cgo !windows pkg-config: opencv4
#cgo CXXFLAGS:   --std=c++11
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/lib -lopencv_gapi -lopencv_videoio -lopencv_video -lopencv_imgcodecs -lopencv_imgproc -lopencv_core
*/
import "C"
