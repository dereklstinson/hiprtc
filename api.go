package hiprtc

/*
#include "hiprtcompile.h"
const void ** voiddptrnull = NULL;
const char* nullchar = NULL;
*/
import "C"
import (
	"errors"
	"runtime"
)

type Program struct {
	p C.gohiprtcProgram
}

type result C.gohiprtcResult

/*CreateProgram srcstring is the contents of the string
srcstring:=`extern "C" __global__ void add(float *A, float *B, float *C,int n){

  int i = blockIdx.x * blockDim.x + threadIdx.x;
    if (i < n) {
       C[tid] = A[i] + B[i];
    }
	}`

	Then you can make a filename := "adder.cpp"

*/
func CreateProgram(srcstring, filename string, headers, includes []string) (p *Program, err error) {
	p = new(Program)
	if headers == nil || includes == nil {
		err = result(C.gohiprtcCreateProgram(&p.p, C.CString(srcstring), C.CString(filename), 0, &C.nullchar, &C.nullchar)).error()
		return p, err
	}
	runtime.SetFinalizer(p, gohiprtcDestroyProgram)

	return nil, errors.New("hiprtc: CreateProgram: Headers and includes not supported yet")
}
func (p *Program) AddNameExpression(nameExpression string) error {
	return result(C.gohiprtcAddNameExpression(p.p, C.CString(nameExpression))).error()

}
func (p *Program) Compile(options []string) error {
	length := (C.int)(len(options))
	opts := make([]*C.char, length)
	for i := range opts {
		opts[i] = C.CString(options[i])
	}

	return result(C.gohiprtcCompileProgram(p.p, length, &opts[0])).error()
}

func gohiprtcDestroyProgram(p *Program) error {
	return result(C.gohiprtcDestroyProgram(&p.p)).error()
}

func (p *Program) GetLoweredName(nameExpression string) (string, error) {
	var char *C.char
	err := result(C.gohiprtcGetLoweredName(p.p, C.CString(nameExpression), &char)).error()
	return C.GoString(char), err

}

/*
gohiprtcResult gohiprtcGetLoweredName(gohiprtcProgram prog,const char* name_expression,const char** lowered_name);
gohiprtcResult gohiprtcGetProgramLog(gohiprtcProgram prog, char* log);
gohiprtcResult gohiprtcGetProgramLogSize(gohiprtcProgram prog,size_t* logSizeRet);
gohiprtcResult gohiprtcGetCode(gohiprtcProgram prog, char* code);
gohiprtcResult gohiprtcGetCodeSize(gohiprtcProgram prog, size_t* codeSizeRet);



*/

func (r result) error() error {
	switch r {
	case (result)(C.HIPRTC_SUCCESS):
		return nil
	case (result)(C.HIPRTC_ERROR_OUT_OF_MEMORY):
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
