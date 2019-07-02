
#include "hiprtcompile.h" 

#include <cstddef>
#include <string>

 

const char* gohiprtcGetErrorString(gohiprtcResult result){
    return hiprtcGetErrorString((hiprtcResult)result);
} 
gohiprtcResult gohiprtcAddNameExpression(gohiprtcProgram prog, const char* name_expression){
    return hiprtcAddNameExpression( prog, name_expression);
}
gohiprtcResult gohiprtcCompileProgram(gohiprtcProgram prog, int numOptions,const char** options){
    return hiprtcCompileProgram( prog,  numOptions,options);
}
gohiprtcResult gohiprtcCreateProgram(gohiprtcProgram* prog, const char* src,const char* name, int numHeaders,const char** headers,const char** includeNames){
    return hiprtcCreateProgram( prog, src, name,  numHeaders,headers,includeNames);
}
gohiprtcResult gohiprtcDestroyProgram(gohiprtcProgram* prog){
    return hiprtcDestroyProgram( prog);
}
gohiprtcResult gohiprtcGetLoweredName(gohiprtcProgram prog,const char* name_expression,const char** lowered_name){
    return hiprtcGetLoweredName( prog,name_expression, lowered_name);
}
gohiprtcResult gohiprtcGetProgramLog(gohiprtcProgram prog, char* log){
    return hiprtcGetProgramLog( prog,  log);
}
gohiprtcResult gohiprtcGetProgramLogSize(gohiprtcProgram prog,size_t* logSizeRet){
    return hiprtcGetProgramLogSize( prog,(std::size_t*) logSizeRet);
}
gohiprtcResult gohiprtcGetCode(gohiprtcProgram prog, char* code){
    return hiprtcGetCode( prog,  code);
}
gohiprtcResult gohiprtcGetCodeSize(gohiprtcProgram prog, size_t* codeSizeRet){
    return hiprtcGetCodeSize( prog, (std::size_t*) codeSizeRet);
}
