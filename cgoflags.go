package hiprtc

/*
#cgo CFLAGS: -D__HIP_PLATFORM_HCC__ -D__HIP_VDI__
#cgo CFLAGS: -I/opt/rocm/hip/include
#cgo CXXFLAGS: --std=c++17
#cgo CXXFLAGS: -I/opt/rocm/hip/include
#cgo LDFLAGS: -L/opt/rocm/hip/lib -lhiprtc -lhip_hcc
*/   
import "C" 
   