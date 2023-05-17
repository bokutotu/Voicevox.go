package voicevox

/*
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/bin
#cgo LDFLAGS: -L${SRCDIR}/bin 
#cgo LDFLAGS: -lvoicevox_core
#include "./bin/voicevox_core.h"
*/
import "C"

type Acceleration C.int

const (
    Auto Acceleration = C.VOICEVOX_ACCELERATION_MODE_AUTO
    CPU Acceleration = C.VOICEVOX_ACCELERATION_MODE_CPU
    GPU Acceleration = C.VOICEVOX_ACCELERATION_MODE_GPU
)
