
#include "hiprtcompile.h" 

#include <cstddef>
#include <string>

 

const char* gohiprtcGetErrorString(gohiprtcResult result){
    return hiprtcGetErrorString((hiprtcResult)result);
} 
gohiprtcResult gohiprtcAddNameExpression(hiprtcProgram prog, const char* name_expression){
    return hiprtcAddNameExpression( prog, name_expression);
}
gohiprtcResult gohiprtcCompileProgram(hiprtcProgram prog, int numOptions,const char** options){
    return hiprtcCompileProgram( prog,  numOptions,options);
}
gohiprtcResult gohiprtcCreateProgram(hiprtcProgram* prog, const char* src,const char* name, int numHeaders,const char** headers,const char** includeNames){
    return hiprtcCreateProgram( prog, src, name,  numHeaders,headers,includeNames);
}
gohiprtcResult gohiprtcDestroyProgram(hiprtcProgram* prog){
    return hiprtcDestroyProgram( prog);
}
gohiprtcResult gohiprtcGetLoweredName(hiprtcProgram prog,const char* name_expression,const char** lowered_name){
    return hiprtcGetLoweredName( prog,name_expression, lowered_name);
}
gohiprtcResult gohiprtcGetProgramLog(hiprtcProgram prog, char* log){
    return hiprtcGetProgramLog( prog,  log);
}
gohiprtcResult gohiprtcGetProgramLogSize(hiprtcProgram prog,size_t* logSizeRet){
    return hiprtcGetProgramLogSize( prog,(std::size_t*) logSizeRet);
}
gohiprtcResult gohiprtcGetCode(hiprtcProgram prog, char* code){
    return hiprtcGetCode( prog,  code);
}
gohiprtcResult gohiprtcGetCodeSize(hiprtcProgram prog, size_t* codeSizeRet){
    return hiprtcGetCodeSize( prog, (std::size_t*) codeSizeRet);
}
