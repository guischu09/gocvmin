#ifndef _OPENCV3_IMGCODECS_H_
#define _OPENCV3_IMGCODECS_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

Mat Image_IMRead(const char* filename, int flags);
Mats Image_IMReadMulti(const char* filename, int flags);
Mats Image_IMReadMulti_WithParams(const char* filename, int start, int count, int flags);
bool Image_IMWrite(const char* filename, Mat img);
bool Image_IMWrite_WithParams(const char* filename, Mat img, IntVector params);
OpenCVResult Image_IMEncode(const char* fileExt, Mat img, void* vector);
OpenCVResult Image_IMEncode_WithParams(const char* fileExt, Mat img, IntVector params, void* vector);
Mat Image_IMDecode(ByteArray buf, int flags);
OpenCVResult Image_IMDecodeIntoMat(ByteArray buf, int flag, Mat dest);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_IMGCODECS_H_
