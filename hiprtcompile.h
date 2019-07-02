#ifndef __HIPRTC_FOR_GO__
#define __HIPRTC_FOR_GO__

#include <stdlib.h>

#ifdef __cplusplus
#define __HIP_PLATFORM_HCC__ 1
#include <hip/hiprtc.h>
extern "C"{
#endif


#ifdef __cplusplus
typedef unsigned int gohiprtcResult;
typedef struct _hiprtcProgram gohiprtcProgram;
#else
typedef unsigned int gohiprtcResult; 
typedef void* gohiprtcProgram;
#endif

#ifdef __cplusplus
}
#endif

#endif //__HIPRTC_FOR_GO__ 