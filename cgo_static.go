//go:build !customenv && opencvstatic && linux

package gocv

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS: --std=c++11
#cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo amd64 LDFLAGS: -O2 -g -static -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_videoio -lopencv_video -lopencv_imgcodecs -lopencv_imgproc -lopencv_core -llibprotobuf -lade -ltbb -littnotify -llibjpeg-turbo -llibwebp -llibtiff -llibopenjp2 -lIlmImf -lquirc -lippiw -lippicv -lpng -lz -lgcc -lstdc++ -lfreetype -lharfbuzz -ldl -lm -lpthread -lrt -lavdevice -lm -latomic -lavfilter -pthread -lm -latomic -lswscale -lm -latomic -lpostproc -lm -latomic -lavformat -lm -latomic -lz -lavcodec -lvpx -lm -lvpx -lm -lvpx -lm -lvpx -lm -pthread -lm -latomic -lz -lx264 -lswresample -lm -latomic -lavutil -pthread -lm -latomic
#cgo arm64 LDFLAGS: -O2 -g -static -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_videoio -lopencv_video -lopencv_imgcodecs -lopencv_imgproc -lopencv_core -llibprotobuf -lade -ltbb -littnotify -llibjpeg-turbo -llibwebp -llibtiff -llibopenjp2 -lIlmImf -lquirc -ltegra_hal -lpng -lz -lgcc -lstdc++ -lfreetype -lharfbuzz -ldl -lm -lpthread -lrt -lavdevice -lm -latomic -lavfilter -pthread -lm -latomic -lswscale -lm -latomic -lpostproc -lm -latomic -lavformat -lm -latomic -lz -lavcodec -lvpx -lm -lvpx -lm -lvpx -lm -lvpx -lm -pthread -lm -latomic -lz -lx264 -lswresample -lm -latomic -lavutil -pthread -lm -latomic
*/
import "C"

// # cgo amd64 LDFLAGS: minimal OpenCV libraries
// # cgo arm64 LDFLAGS: minimal OpenCV libraries
