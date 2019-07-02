#ifndef __HIPRTC_FOR_GO__
#define __HIPRTC_FOR_GO__

#include <stdlib.h>

#ifdef __cplusplus
#define __HIP_PLATFORM_HCC__ 1
#define HIPRTC_GET_TYPE_NAME 1
#include <hip/hiprtc.h>

extern "C"{
#endif
//typedef struct _hiprtcProgram *gohiprtcProgram;
//typedef enum hiprtcResult gohiprtcResult;
 
#ifdef __cplusplus
typedef unsigned int gohiprtcResult;
typedef struct _hiprtcProgram* gohiprtcProgram;
#else
typedef enum hiprtcResultx {
    HIPRTC_SUCCESS = 0,
    HIPRTC_ERROR_OUT_OF_MEMORY = 1,
    HIPRTC_ERROR_PROGRAM_CREATION_FAILURE = 2,
    HIPRTC_ERROR_INVALID_INPUT = 3,
    HIPRTC_ERROR_INVALID_PROGRAM = 4, 
    HIPRTC_ERROR_INVALID_OPTION = 5,
    HIPRTC_ERROR_COMPILATION = 6,
    HIPRTC_ERROR_BUILTIN_OPERATION_FAILURE = 7,
    HIPRTC_ERROR_NO_NAME_EXPRESSIONS_AFTER_COMPILATION = 8,
    HIPRTC_ERROR_NO_LOWERED_NAMES_BEFORE_COMPILATION = 9,
    HIPRTC_ERROR_NAME_EXPRESSION_NOT_VALID = 10,
    HIPRTC_ERROR_INTERNAL_ERROR = 11
} gohiprtcResult; 
typedef void* gohiprtcProgram;
#endif

const char* gohiprtcGetErrorString(gohiprtcResult result);
gohiprtcResult gohiprtcAddNameExpression(gohiprtcProgram prog, const char* name_expression);
 
gohiprtcResult gohiprtcCompileProgram(gohiprtcProgram prog, int numOptions,const char** options);
gohiprtcResult gohiprtcCreateProgram(gohiprtcProgram* prog, const char* src,const char* name, int numHeaders,const char** headers,const char** includeNames);
gohiprtcResult gohiprtcDestroyProgram(gohiprtcProgram* prog);
gohiprtcResult gohiprtcGetLoweredName(gohiprtcProgram prog,const char* name_expression,const char** lowered_name);
gohiprtcResult gohiprtcGetProgramLog(gohiprtcProgram prog, char* log);
gohiprtcResult gohiprtcGetProgramLogSize(gohiprtcProgram prog,size_t* logSizeRet);
gohiprtcResult gohiprtcGetCode(gohiprtcProgram prog, char* code);
gohiprtcResult gohiprtcGetCodeSize(gohiprtcProgram prog, size_t* codeSizeRet);

#ifdef __cplusplus
}
#endif

#endif //__HIPRTC_FOR_GO__ 