package hiprtc
/*
#include "hiprtcompile.h"
const void* nullptr = NULL;
*/
import "C"
import "errors"

type Program struct{
	p C.gohiprtcProgram
}     

type result C.gohiprtcResult 

func CreateProgram(src, name string, headers, includes []string )(p *Program, err error){
	p=new(Program)

err=result(C.gohiprtcCreateProgram(&p.p, C.CString(src), C.CString(name) ,0,&(*C.char)(C.nullptr),&((*C.char)(C.nullptr)))).error()
return p,err
}
/*
gohiprtcResult gohiprtcAddNameExpression(gohiprtcProgram prog, const char* name_expression);
 
gohiprtcResult gohiprtcCompileProgram(gohiprtcProgram prog, int numOptions,const char** options);
gohiprtcResult gohiprtcCreateProgram(gohiprtcProgram* prog, const char* src,const char* name, int numHeaders,const char** headers,const char** includeNames);
gohiprtcResult gohiprtcDestroyProgram(gohiprtcProgram* prog);
gohiprtcResult gohiprtcGetLoweredName(gohiprtcProgram prog,const char* name_expression,const char** lowered_name);
gohiprtcResult gohiprtcGetProgramLog(gohiprtcProgram prog, char* log);
gohiprtcResult gohiprtcGetProgramLogSize(gohiprtcProgram prog,size_t* logSizeRet);
gohiprtcResult gohiprtcGetCode(gohiprtcProgram prog, char* code);
gohiprtcResult gohiprtcGetCodeSize(gohiprtcProgram prog, size_t* codeSizeRet);



*/





func (r result)error()error{
	switch r{
	case (result)(C.HIPRTC_SUCCESS):
		return nil
	case (result)(C.HIPRTC_ERROR_OUT_OF_MEMORY ):
		return errors.New("HIPRTC_ERROR_OUT_OF_MEMORY ")
	case (result)(C.HIPRTC_ERROR_PROGRAM_CREATION_FAILURE):
		return errors.New("HIPRTC_ERROR_PROGRAM_CREATION_FAILURE")
	case (result)(C.HIPRTC_ERROR_INVALID_INPUT):
		return errors.New("HIPRTC_ERROR_INVALID_INPUT")
	case (result)(C.HIPRTC_ERROR_INVALID_PROGRAM): 
		return errors.New("HIPRTC_ERROR_INVALID_PROGRAM")
	case (result)(C.HIPRTC_ERROR_INVALID_OPTION):
		return errors.New("HIPRTC_ERROR_INVALID_OPTION")
	case (result)(C.HIPRTC_ERROR_COMPILATION):
		return errors.New("HIPRTC_ERROR_COMPILATION")
	case (result)(C.HIPRTC_ERROR_BUILTIN_OPERATION_FAILURE):
		return errors.New("HIPRTC_ERROR_BUILTIN_OPERATION_FAILURE")
	case (result)(C.HIPRTC_ERROR_NO_NAME_EXPRESSIONS_AFTER_COMPILATION):
		return errors.New("HIPRTC_ERROR_NO_NAME_EXPRESSIONS_AFTER_COMPILATION")
	case (result)(C.HIPRTC_ERROR_NO_LOWERED_NAMES_BEFORE_COMPILATION):
		return errors.New("HIPRTC_ERROR_NO_LOWERED_NAMES_BEFORE_COMPILATION")
	case (result)(C.HIPRTC_ERROR_NAME_EXPRESSION_NOT_VALID):
		return errors.New("HIPRTC_ERROR_NAME_EXPRESSION_NOT_VALID")
	case (result)(C.HIPRTC_ERROR_INTERNAL_ERROR):
		return errors.New("HIPRTC_ERROR_INTERNAL_ERROR")
	}
	return errors.New("Undefined error caught on go side")
}
 